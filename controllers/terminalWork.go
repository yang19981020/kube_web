package controllers

import beego "github.com/beego/beego/v2/server/web"

type TerminalController struct {
	beego.Controller
}

func (self *TerminalController) Get() {
	//环境信息
	self.TplName = "terminal/index.html"
}