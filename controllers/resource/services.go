package resource

import beego "github.com/beego/beego/v2/server/web"

type ServicesController struct {
	beego.Controller
}

func (c *ServicesController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/services/index.html"
}
