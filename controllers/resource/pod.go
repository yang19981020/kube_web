package resource

import (
	"context"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kube_web/common"
	"kube_web/common/response"
	"kube_web/models"
	m "kube_web/models/resource"
	s "kube_web/services/resource"
)

type PodController struct {
	beego.Controller
}

func (c *PodController) PodListApi() {
	get_str := c.GetString("namespace", "default")
	sClient := common.K8sClient
	pod_list, _ := sClient.CoreV1().Pods(fmt.Sprint(get_str)).List(context.Background(),v1.ListOptions{})
	json := response.Json(200, "ok", pod_list)
	c.Data["json"] = json
	c.ServeJSON()
}

func (c *PodController) Get() {
	get_namespace := c.GetString("namespace", "default")
	get_podname := c.GetString("podName", "")
	client := s.New()
	pod, err := client.GetPod(get_namespace, get_podname)
    if err != nil {
		json := models.NewError(400, models.ErrGetPOD.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", pod)
	c.ServeJSON()
}

func (c *PodController) Post()() {
	get_namespace := c.GetString("namespace", "default")
	var podForm = &m.PodParams{}
	c.ParseForm(podForm)
	client := s.New()
	pod, err := client.CreatePod(get_namespace, podForm)
	if err != nil {
		json := models.NewError(400, models.ErrCreatePOD.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok", pod)
}

func (c *PodController) Delete() {
	get_namespace := c.GetString("namespace", "default")
	get_podname := c.GetString("podName", "")
	client := s.New()
	err := client.DeletePod(get_namespace, get_podname)
	if err != nil {
		json := models.NewError(400, models.ErrDelPOD.Error())
		c.Data["json"] = json
		c.ServeJSON()
	}
	c.Data["json"] = response.Json(200,"ok")
}







