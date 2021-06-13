package resource

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utils2 "kube_web/common"
	"kube_web/common/response"
)

type NodeListController struct {
	beego.Controller
}

func (c *NodeListController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/nodeList/index.html"
}

func (c *NodeListController) NodeListApi(){
	client := utils2.K8sClient
	node_list, _ := client.CoreV1().Nodes().List(context.Background(),v1.ListOptions{})
	c.Data["json"] = response.Json(200,"ok",node_list)
	c.ServeJSON()
}
