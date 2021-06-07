package resource

import beego "github.com/beego/beego/v2/server/web"

type ServiceAccountsController struct {
	beego.Controller
}

func (c *ServiceAccountsController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/serviceAccounts/index.html"
}
