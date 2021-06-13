package resource

import (
	"context"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube_web/common"
	"kube_web/common/response"
	"kube_web/models"
	s "kube_web/services/resource"
)

type NamespaceController struct {
	beego.Controller
}

func (c *NamespaceController) NamespaceListApi() {
	client := common.K8sClient
	namespaceList,_ := client.CoreV1().Namespaces().List(context.Background(),v1.ListOptions{})
	json := response.Json(200, "ok", namespaceList)
	c.Data["json"] = json
	c.ServeJSON()
}
func (c *NamespaceController) Get() {
	get_namespace := c.GetString("namespace", "default")
	client := s.New()
	namespace , err := client.GetNamespace(get_namespace)
	if err != nil {
		json := models.NewError(400, models.ErrGetNamespace.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", namespace)
}
func (c *NamespaceController) Post() {
	get_namespace := c.GetString("namespace", "default")
	client := s.New()
	namespace , err := client.CreateNamespace(get_namespace)
	if err != nil {
		json := models.NewError(400, models.ErrCreateNamespace.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", namespace)
	c.ServeJSON()
}
func (c *NamespaceController) Delete() {
	get_namespace := c.GetString("namespace", "default")
	client := s.New()
	err := client.DeleteNamespace(get_namespace)
	if err != nil {
		json := models.NewError(400, models.ErrDeleteNamespace.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok")
	c.ServeJSON()
}



