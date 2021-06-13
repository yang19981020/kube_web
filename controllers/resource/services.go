package resource

import (
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/common/response"
	"kube_web/models"
	m "kube_web/models/resource"
	s "kube_web/services/resource"
)

type ServicesController struct {
	beego.Controller
}

func (c *ServicesController) ListServiceApi() {
	get_namespace := c.GetString("namespace", "default")
	client := s.New()
	service_list , err := client.ListService(get_namespace)
	if err != nil {
		json := models.NewError(400, models.ErrListService.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", service_list)
	c.ServeJSON()
}
func (c *ServicesController) Get() {
	get_namespace := c.GetString("namespace", "default")
	get_name := c.GetString("name", "")
	client := s.New()
	service , err := client.GetService(get_namespace,get_name)
	if err != nil {
		json := models.NewError(400, models.ErrGetService.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", service)
	c.ServeJSON()
}
func (c *ServicesController) Post() {
	get_namespace := c.GetString("namespace", "default")
	var svcForm = &m.ServiceParams{}
	err := c.ParseForm(svcForm)
	if err != nil {
		json := models.NewError(400, models.ErrParseFormSVC.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	client := s.New()
	svc , err := client.CreateService(get_namespace, svcForm)
	if err != nil {
		json := models.NewError(400, models.ErrCreateService.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", svc)
	c.ServeJSON()
}
func (c *ServicesController) Delete() {
	get_namespace := c.GetString("namespace", "default")
	get_name := c.GetString("name", "")
	client := s.New()
	err := client.DeleteService(get_namespace, get_name)
	if err != nil {
		json := models.NewError(400, models.ErrDeleteService.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok")
	c.ServeJSON()
}
