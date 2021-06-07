package controllers

import beego "github.com/beego/beego/v2/server/web"

type LoginController struct{
	beego.Controller
}

func (c *LoginController) Get(){
	c.TplName = "login/index.html"
}