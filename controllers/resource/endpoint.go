package resource

import beego "github.com/beego/beego/v2/server/web"

type EndpointController struct {
	beego.Controller
}

func (c *EndpointController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/endpoint/index.html"
}
