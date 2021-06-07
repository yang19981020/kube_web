package resource

import beego "github.com/beego/beego/v2/server/web"

type ClusterRoleBindingController struct {
	beego.Controller
}

func (c *ClusterRoleBindingController) Get() {
	c.Layout = "index.html"
	c.TplName = "body/clusterRoleBinding/index.html"
}
