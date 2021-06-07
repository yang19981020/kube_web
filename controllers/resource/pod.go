package resource

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube_web/utils"
	"kube_web/utils/response"
)

type PodController struct {
	beego.Controller
}

func (c *PodController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/pod/index.html"
}

func (c *PodController) PodListApi() {
	get_str := c.GetString("namespace", "default")
	sClient := utils.K8sClient
	pod_list, _ := sClient.CoreV1().Pods(fmt.Sprint(get_str)).List(context.Background(),v1.ListOptions{})
	json := response.Json(200, "ok", pod_list)
	c.Data["json"] = json
	c.ServeJSON()
}
