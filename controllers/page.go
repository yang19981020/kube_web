package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type PageController struct{
	beego.Controller
}

func (c *PageController) Get(){
	c.TplName = "index.html"
}