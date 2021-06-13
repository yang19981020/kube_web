package main

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	_ "kube_web/routers"
	"time"
)

func dbinit()  {
	runmode,_ := beego.AppConfig.String("runmode")
	isDev := (runmode == "dev")
	registDatabase()
	if isDev {
		orm.Debug = isDev
		// 非强制模式下自动建表
		err := orm.RunSyncdb("default", false, isDev)
		if err != nil {
			logs.Informational("[orm] Create table err : ", err)
		}
	}
}
func registDatabase()  {
	//初始化数据库
	dbUser, _ := beego.AppConfig.String("mysqluser")
	dbPass, _ := beego.AppConfig.String("mysqlpass")
	dbName, _ := beego.AppConfig.String("mysqldb")
	dbHost, _ := beego.AppConfig.String("mysqlhost")
	dbPort, _ := beego.AppConfig.String("mysqlport")
	maxIdleConn, _ := beego.AppConfig.Int("db_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db_max_open_open")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql",
		dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=true&loc=Asia%2FShanghai",
	)
	orm.MaxIdleConnections(maxIdleConn)
	orm.MaxOpenConnections(maxOpenConn)
	orm.DefaultTimeLoc = time.UTC
}
func loginit()  {
	//日志
	logs.Async()
	level, _ := beego.AppConfig.Int("logLevel")
	logs.SetLevel(level)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"./logs/kube_web.log",
	"level":6,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)
}


func init() {
	dbinit()
	loginit()
}

func main() {
	beego.Run()
}
