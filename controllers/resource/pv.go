package resource

import beego "github.com/beego/beego/v2/server/web"

type PvController struct {
	beego.Controller
}

func (c *PvController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/pv/index.html"
}
