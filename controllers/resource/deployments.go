package resource

import (
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/common/response"
	"kube_web/models"
	m "kube_web/models/resource"
	s "kube_web/services/resource"
)

type DeploymentsController struct {
	beego.Controller
}

func (d *DeploymentsController) GetAll() {
	get_namespace := d.GetString("namespace", "default")
	client := s.New()
	list, err := client.ListDeployment(get_namespace)
	if err != nil {
		json := models.NewError(400, models.ErrListDeployment.Error())
		d.Data["json"] = json
		d.ServeJSON()
	}
	d.Data["json"] = response.Json(200,"ok", list)
}
func (d *DeploymentsController) Get() {
	get_namespace := d.GetString("namespace", "default")
	get_deployment := d.GetString("deployment", "")
	client := s.New()
	dep , err := client.GetDeployment(get_namespace,get_deployment)
	if err != nil {
		json := models.NewError(400, models.ErrGetDeployment.Error())
		d.Data["json"] = json
		d.ServeJSON()
	}
	d.Data["json"] = response.Json(200,"ok", dep)
}

func (d *DeploymentsController) Delete() {
	get_namespace := d.GetString("namespace", "default")
	get_deployment := d.GetString("deployment", "")
	client := s.New()
	err := client.DeleteDeployment(get_namespace,get_deployment)
	if err != nil {
		json := models.NewError(400, models.ErrDeleteDeployment.Error())
		d.Data["json"] = json
		d.ServeJSON()
	}
	d.Data["json"] = response.Json(200,"ok")
}

func (d *DeploymentsController) Post()() {
	get_namespace := d.GetString("namespace", "default")
	get_update,_ := d.GetBool("update", false)
	var depForm = &m.DeploymentParams{}
	err := d.ParseForm(depForm)
	if err != nil {
		json := models.NewError(400, models.ErrDeploymentParams.Error())
		d.Data["json"] = json
		d.ServeJSON()
	}
	client := s.New()
	deployment, err := client.ApplyDeployment(get_namespace, depForm,get_update )
	if err != nil {
		json := models.NewError(400, models.ErrApplyDeployment.Error())
		d.Data["json"] = json
		d.ServeJSON()
	}
	d.Data["json"] = response.Json(200,"ok",deployment)
}