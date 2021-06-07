package resource

import beego "github.com/beego/beego/v2/server/web"

type DeploymentsController struct {
	beego.Controller
}

func (c *DeploymentsController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/deployments/index.html"
}