package controllers

import (
	"github.com/beego/beego/v2/adapter/logs"
	beego "github.com/beego/beego/v2/server/web"
	jwt "kube_web/common/jwt"
	"kube_web/models"
	"strings"
	"time"
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	UserName       string
}
var singinKey,_ = beego.AppConfig.String("jwt_token")

// 初始化函数
func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName)
	this.actionName = strings.ToLower(actionName)

	// pass login
	if controllerName == "Login" && actionName == "Login" || actionName == "Captcha" {
		return
	}

	if this.Ctx.Input.Header("Authorization") == "" {
		this.ResponseJson(401, "签名为空", nil)
		this.StopRun()
	}
	token := strings.TrimSpace(strings.TrimLeft(this.Ctx.Input.Header("Authorization"), "Bearer"))
	if token == "" {
		logs.Error("unknown token:", token)
		this.ResponseJson(401, "未知签名", nil)
		this.StopRun()
	}
	_, err := this.CheckToken(token)
	if err != nil {
		logs.Error(err)
		this.ResponseJson(401, "无效的签名", nil)
		this.StopRun()
	}

}

// 响应json
func (this *BaseController) ResponseJson(code int, msg string, data interface{}) {
	ret := map[string]interface{}{"code": code, "msg": msg}
	if data != nil {
		ret["data"] = data
	} else {
		ret["data"] = map[string]interface{}{}
	}

	this.Data["json"] = ret
	this.ServeJSON()
}


func (this *BaseController) GetToken(m *models.SysUser) (token string, err error) {
	token, err = jwt.GenerateToken(m, time.Hour*24)
	return
}

func (this *BaseController) CheckToken(tokenStr string) (user *models.SysUser,err error) {
	user, err = jwt.ValidateToken(tokenStr)
	return
}

type Result struct {
	Data interface{} `json:"data"`
	Msg string       `json:"msg"`
	Status int       `json:"status"`
}

func (c *BaseController) Ok(data interface{})  {
	c.Data["json"] = SuccessData(data)
	c.ServeJSON()
}

func (c *BaseController) Fail(msg string,status int)  {
	c.Data["json"] = ErrMsg(msg,status)
	c.ServeJSON()
}

func ErrMsg(msg string,status ...int) Result {
	var r Result
	if len(status) > 0 {
		r.Status = status[0]
	}else{
		r.Status = 400
	}
	r.Msg = msg
	r.Data = nil

	return r
}

func ErrData(msg error,status ...int) Result {
	var r Result
	if len(status) > 0 {
		r.Status = status[0]
	} else {
		r.Status = 500000
	}
	r.Msg = msg.Error()
	r.Data = nil

	return r
}

func SuccessData(data interface{}) Result {
	var r Result

	r.Status = 200
	r.Msg = "ok"
	r.Data = data

	return r
}