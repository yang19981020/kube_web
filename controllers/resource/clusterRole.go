package resource

import beego "github.com/beego/beego/v2/server/web"

type ClusterRoleController struct {
	beego.Controller
}

func (c *ClusterRoleController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/clusterRole/index.html"
}