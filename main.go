package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "kube_web/routers"
)

func main() {
	beego.Run()
}

