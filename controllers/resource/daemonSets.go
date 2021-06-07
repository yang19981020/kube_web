package resource

import beego "github.com/beego/beego/v2/server/web"

type DaemonSetsController struct {
	beego.Controller
}

func (c *DaemonSetsController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/daemonSets/index.html"
}