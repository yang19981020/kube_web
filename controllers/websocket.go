package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gorilla/websocket"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"kube_web/common"
	"net/http"
	"sync"
)

// http升级websocket协议的配置
var wsUpgrader = websocket.Upgrader{
	// 允许所有CORS跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket消息
type WsMessage struct {
	MessageType int
	Data        []byte
}

// 封装websocket连接
type WsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *WsMessage // 读取队列
	outChan  chan *WsMessage // 发送队列

	mutex     sync.Mutex // 避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

// 读取协程
func (wsConn *WsConnection) wsReadLoop() {
	var (
		msgType int
		data    []byte
		msg     *WsMessage
		err     error
	)
	for {
		// 读一个message
		if msgType, data, err = wsConn.wsSocket.ReadMessage(); err != nil {
			goto ERROR
		}
		msg = &WsMessage{
			msgType,
			data,
		}
		// 放入请求队列
		select {
		case wsConn.inChan <- msg:
		case <-wsConn.closeChan:
			goto CLOSED
		}
	}
ERROR:
	logs.Debug("读取协程报错",err.Error())
	wsConn.WsClose()
CLOSED:
}

// 发送协程
func (wsConn *WsConnection) wsWriteLoop() {
	var (
		msg *WsMessage
		err error
	)
	for {
		select {
		// 取一个应答
		case msg = <-wsConn.outChan:
			// 写给websocket
			if err = wsConn.wsSocket.WriteMessage(msg.MessageType, msg.Data); err != nil {
				goto ERROR
			}
		case <-wsConn.closeChan:
			goto CLOSED
		}
	}
ERROR:
	logs.Debug("发送协程报错",err.Error())
	wsConn.WsClose()
CLOSED:
}

/************** 并发安全 API **************/

func InitWebsocket(resp http.ResponseWriter, req *http.Request) (wsConn *WsConnection, err error) {
	var (
		wsSocket *websocket.Conn
	)
	// 应答客户端告知升级连接为websocket
	if wsSocket, err = wsUpgrader.Upgrade(resp, req, nil); err != nil {
		return
	}
	wsConn = &WsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *WsMessage, 1000),
		outChan:   make(chan *WsMessage, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()

	return
}

// 发送消息
func (wsConn *WsConnection) WsWrite(messageType int, data []byte) (err error) {
	select {
	case wsConn.outChan <- &WsMessage{messageType, data}:
	case <-wsConn.closeChan:
		err = errors.New("websocket closed")
	}
	return
}

// 读取消息
func (wsConn *WsConnection) WsRead() (msg *WsMessage, err error) {
	select {
	case msg = <-wsConn.inChan:
		return
	case <-wsConn.closeChan:
		err = errors.New("websocket closed")
	}
	return
}

// 关闭连接
func (wsConn *WsConnection) WsClose() {
	wsConn.wsSocket.Close()

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}

// ssh流式处理器
type streamHandler struct {
	wsConn      *WsConnection
	resizeEvent chan remotecommand.TerminalSize
}

// web终端发来的包
type xtermMessage struct {
	MsgType string `json:"type"`  // 类型:resize客户端调整终端, input客户端输入
	Input   string `json:"input"` // msgtype=input情况下使用
	Rows    uint16 `json:"rows"`  // msgtype=resize情况下使用
	Cols    uint16 `json:"cols"`  // msgtype=resize情况下使用
}

// executor回调获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	ret := <-handler.resizeEvent
	size = &ret
	return
}

// executor回调读取web端的输入
func (handler *streamHandler) Read(p []byte) (size int, err error) {
	var (
		msg      *WsMessage
		xtermMsg xtermMessage
	)
	// 读web发来的输入
	if msg, err = handler.wsConn.WsRead(); err != nil {
		return
	}
	// 解析客户端请求
	if err = json.Unmarshal(msg.Data, &xtermMsg); err != nil {
		return
	}
	//web ssh调整了终端大小
	if xtermMsg.MsgType == "resize" {
		handler.resizeEvent <- remotecommand.TerminalSize{Width: xtermMsg.Cols, Height: xtermMsg.Rows}
	} else if xtermMsg.MsgType == "input" { // web ssh终端输入了字符
		// copy到p数组中
		size = len(xtermMsg.Input)
		copy(p, xtermMsg.Input)
	}
	return
}

// executor回调向web端输出
func (handler *streamHandler) Write(p []byte) (size int, err error) {
	var (
		copyData []byte
	)
	// 产生副本
	copyData = make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	err = handler.wsConn.WsWrite(websocket.TextMessage, copyData)
	return
}

//http接口实现
func (self *WsConnection) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// ws = new WebSocket("ws://localhost:8080/resource/websocket?" + "podNs=" + podNs + "&podName=" + podName + "&containerName=" + containerName );
	var (
		restConf      *rest.Config
		clientset     *kubernetes.Clientset
		sshReq        *rest.Request
		podName       string
		podNs         string
		containerName string
		shell         string
		executor      remotecommand.Executor
		wsConn        *WsConnection
		handler       *streamHandler
		err           error
	)
	// 解析GET参数
	if err = req.ParseForm(); err != nil {
		return
	}
	podNs = req.Form.Get("podNs")
	podName = req.Form.Get("podName")
	containerName = req.Form.Get("containerName")
	shell = req.Form.Get("shell")
	// 得到websocket长连接
	if wsConn, err = InitWebsocket(resp, req); err != nil {
		return
	}
	// 获取k8s rest conf配置
	restConf = common.K8sConfig
	// 获取k8s client配置
	clientset = common.K8sClient
	// https://172.18.11.25:6443/api/v1/namespaces/default/pods/nginx-deployment-5cbd8757f-d5qvx/exec?command=sh&container=nginx&stderr=true&stdin=true&stdout=true&tty=true
	if shell == "" { shell = "bash"}
	sshReq = clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(podNs).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{shell},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)
	// 创建到容器的连接
	if executor, err = remotecommand.NewSPDYExecutor(restConf, "POST", sshReq.URL()); err != nil {
		logs.Debug("创建到容器的连接", err.Error())
		wsConn.WsClose()
	}
	// 配置与容器之间的数据流处理回调
	handler = &streamHandler{wsConn: wsConn, resizeEvent: make(chan remotecommand.TerminalSize)}
	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		TerminalSizeQueue: handler,
		Tty:               true,
	}); err != nil {
		logs.Debug("配置与容器之间的数据流处理回调", err.Error())
		wsConn.WsClose()
	}
	return
}
