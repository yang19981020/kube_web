package resource

import beego "github.com/beego/beego/v2/server/web"

type ConfigMapController struct {
	beego.Controller
}

func (c *ConfigMapController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/configMap/index.html"
}
