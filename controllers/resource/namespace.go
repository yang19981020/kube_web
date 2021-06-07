package resource

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube_web/utils"
	"kube_web/utils/response"
)

type NamespaceController struct {
	beego.Controller
}

func (c *NamespaceController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/namespace/index.html"
}

func (c *NamespaceController) NamespaceListApi() {
	client := utils.K8sClient
	namespaceList,_ := client.CoreV1().Namespaces().List(context.Background(),v1.ListOptions{})
	json := response.Json(200, "ok", namespaceList)
	c.Data["json"] = json
	c.ServeJSON()
}