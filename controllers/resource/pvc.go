package resource

import beego "github.com/beego/beego/v2/server/web"

type PvcController struct {
	beego.Controller
}

func (c *PvcController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/pvc/index.html"
}
