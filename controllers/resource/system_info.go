package resource

import (
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/utils/response"
	"kube_web/utils/system"
)

type SystemInfoController struct {
	beego.Controller
}

func (u *SystemInfoController) Get() {
	Map := map[string]interface{}{
		"mem": system.GetMemPercent(),
		"cpu": system.GetCpuPercent(),
	}
	json := response.Json(200, "OK", Map)
	u.Data["json"] = json
	u.ServeJSON()
}
