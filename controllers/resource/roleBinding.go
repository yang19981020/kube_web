package resource

import beego "github.com/beego/beego/v2/server/web"

type RoleBindingController struct {
	beego.Controller
}

func (c *RoleBindingController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/roleBinding/index.html"
}
