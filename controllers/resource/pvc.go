package resource

import (
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/models"
	m "kube_web/models/resource"
	s "kube_web/services/resource"
	"kube_web/utils/response"
)

type PvcController struct {
	beego.Controller
}

func (c *PvcController) Post() {
	get_namespace := c.GetString("namespace", "default")
	var pvcForm = &m.PVCParams{}
	err := c.ParseForm(pvcForm)
	if err!= nil {

	}
	client := s.New()
	pvc, err := client.CreatePVC(get_namespace, pvcForm)
	if err != nil {
		json := models.NewError(400, models.ErrCreatePVC.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", pvc)
}