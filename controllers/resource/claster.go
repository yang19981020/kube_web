package resource

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gogf/gf/frame/g"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utils "kube_web/utils"
	"kube_web/utils/response"
)

type ClasterController struct {
	beego.Controller
}

func (c *ClasterController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/cluster/index.html"
}

func (c *ClasterController) ClusterInfo(){
	client := utils.K8sClient
	node_list, _ := client.CoreV1().Nodes().List(context.Background(),v1.ListOptions{})
	servicelist, _ := client.CoreV1().Services("").List(context.Background(),v1.ListOptions{})
	pod_list, _ := client.CoreV1().Pods("").List(context.Background(),v1.ListOptions{})
	deplist,_ := client.AppsV1().Deployments("").List(context.Background(),v1.ListOptions{})
	daeonlist, _ := client.AppsV1().DaemonSets("").List(context.Background(),v1.ListOptions{})
	statefulsetlist, _ := client.AppsV1().StatefulSets("").List(context.Background(),v1.ListOptions{})

	c.Data["json"] = response.Json(200,"ok",g.Map{
		"node": len(node_list.Items),
		"service": len(servicelist.Items),
		"deployment": len(deplist.Items),
		"statefulset": len(statefulsetlist.Items),
		"pod": len(pod_list.Items),
		"daemonset": len(daeonlist.Items),
	})
	c.ServeJSON()
}
