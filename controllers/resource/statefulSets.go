package resource

import beego "github.com/beego/beego/v2/server/web"

type StatefulSetsController struct {
	beego.Controller
}

func (c *StatefulSetsController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/statefulSets/index.html"
}
