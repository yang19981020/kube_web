package resource

import (
	beego "github.com/beego/beego/v2/server/web"
)

type StatefulSetsController struct {
	beego.Controller
}

//func (c *StatefulSetsController) Get() {
//	get_namespace := c.GetString("namespace", "default")
//	sClient := utils.K8sClient
//	pod_list, _ := sClient.AppsV1().StatefulSets(get_namespace).List(c.,v1.ListOptions{})
//	json := response.Json(200, "ok", pod_list)
//	c.Data["json"] = json
//	c.ServeJSON()
//}