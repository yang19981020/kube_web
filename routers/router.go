package routers

import (
	"kube_web/controllers"
	"kube_web/controllers/resource"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "token"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/resource",
		beego.NSRouter("/system_info", &resource.SystemInfoController{}),
		beego.NSRouter("/cluster_info", &resource.ClasterController{},"get:ClusterInfo"),
		beego.NSRouter("/namespace_list", &resource.NamespaceController{},"get:NamespaceListApi"),
		beego.NSRouter("/pod_list", &resource.PodController{},"get:PodListApi"),
		beego.NSRouter("/node_list", &resource.NodeListController{},"get:NodeListApi"),
	    beego.NSRouter("/pod", &resource.PodController{}),
	    beego.NSRouter("/deployment", &resource.DeploymentsController{}),
	    beego.NSRouter("/namespace", &resource.NamespaceController{}),
	    beego.NSRouter("/service", &resource.ServicesController{}),
	    beego.NSRouter("/pvc", &resource.PvcController{}),
	)
	beego.Handler("/resource/websocket",&controllers.WsConnection{},true)
	ns1 := beego.NewNamespace("/v1",
		beego.NSRouter("/picCode", &controllers.Login{},"get:Captcha"),
		beego.NSRouter("/login", &controllers.Login{},"post:Login"),
		)
	beego.AddNamespace(ns,ns1)
}
