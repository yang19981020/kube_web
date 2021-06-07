package resource

import beego "github.com/beego/beego/v2/server/web"

type RoleController struct {
	beego.Controller
}

func (c *RoleController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/role/index.html"
}
