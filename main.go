package main

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/adapter/orm"
	beego "github.com/beego/beego/v2/server/web"
	"kube_web/models"
	_ "kube_web/routers"
)

func init(){
	orm.Debug = true

	if err1 := orm.RegisterDriver("mysql", orm.DRMySQL); err1 != nil {
		logs.Error(err1.Error())
	}
	orm.RegisterModel(new(models.User))

	if err2 := orm.RegisterDataBase("default","mysql","root:12345678@tcp(127.0.0.1:3306)/testdb");err2 != nil {
		logs.Error(err2.Error())
		panic(err2.Error())
	}
	fmt.Println("Connected to the database")
	orm.RunSyncdb("default", false, true)
}

func main() {


	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

