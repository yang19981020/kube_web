package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	beegoormadapter "github.com/casbin/beego-orm-adapter/v3"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	_ "kube_web/routers"
)

func init() {
	// 参数4(可选)  设置最大空闲连接
	// 参数5(可选)  设置最大数据库连接 (go >= 1.2)
	maxIdle := 30
	maxConn := 30
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123123Aa..@tcp(114.67.110.204:3306)/testdb?charset=utf8",orm.MaxIdleConnections(maxIdle), orm.MaxOpenConnections(maxConn))
	orm.Debug = true
	//file, err := os.Create("orm.txt")
	//if err != nil {
	//	logs.Warn("open file fail")
	//	os.Exit(1)
	//}
	//orm.DebugLog = orm.NewLog(file)

	a, err := beegoormadapter.NewAdapter("default1", "mysql", "root:123123Aa..@tcp(114.67.110.204:3306)/testdb?charset=utf8") // Your driver and data source.
	if err != nil {
		logs.Warn(err)
	}
	e, err := casbin.NewEnforcer("conf/rbac_model.conf", a)
	if err != nil {
		logs.Warn(err)
	}
	// Check the permission.
	enforce, err := e.Enforce("alice", "data1", "read") // 用户 资源 权限
	fmt.Println(enforce,err)
	policy, err := e.AddPolicy("alice", "data1", "read")
	fmt.Println(policy,err)
	//removePolicy, err := e.RemovePolicy("alice", "data1", "read")
	//fmt.Println(removePolicy,err)

	// Save the policy back to DB.
	e.SavePolicy()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
