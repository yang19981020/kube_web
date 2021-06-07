package resource

import beego "github.com/beego/beego/v2/server/web"

type SecretsController struct {
	beego.Controller
}

func (c *SecretsController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/secrets/index.html"
}

