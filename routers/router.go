package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"kube_web/controllers"
	"kube_web/controllers/resource"
)

func init() {
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
	beego.Router("/", &resource.ClasterController{})
	beego.Router("/pod", &resource.PodController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/node_list", &resource.NodeListController{})
	beego.Router("/cluster_role", &resource.ClusterRoleController{})
	beego.Router("/cluster_role_binding", &resource.ClusterRoleBindingController{})
	beego.Router("/config_map", &resource.ConfigMapController{})
	beego.Router("/daemon_set", &resource.DaemonSetsController{})
	beego.Router("/deployment", &resource.DeploymentsController{})
	beego.Router("/namespace", &resource.NamespaceController{})
	beego.Router("/pv", &resource.PvController{})
	beego.Router("/pvc", &resource.PvcController{})
	beego.Router("/role", &resource.RoleController{})
	beego.Router("/role_binding", &resource.RoleBindingController{})
	beego.Router("/secrets", &resource.SecretsController{})
	beego.Router("/service", &resource.ServicesController{})
	beego.Router("/service_account", &resource.ServiceAccountsController{})
	beego.Router("/endpoint", &resource.EndpointController{})
	beego.Router("/stateful_set", &resource.StatefulSetsController{})
	beego.Router("/deployment_edit", &resource.DeploymentsController{})
	beego.Router("/deployment_create", &resource.DeploymentsController{})
	// http://127.0.0.1:8080/websocket?namespace=default&pod=nginx-01-1&container=nginx-01
	beego.Router("/websocket",&controllers.TerminalController{})
	//beego.Handler("/resource/websocket",&controllers.WebSocketStruct{},true)
	beego.Handler("/resource/websocket",&controllers.WsConnection{},true)

	ns := beego.NewNamespace("/resource",
		beego.NSRouter("/system_info", &resource.SystemInfoController{}),
		beego.NSRouter("/cluster_info", &resource.ClasterController{},"get:ClusterInfo"),
		beego.NSRouter("/namespace_list", &resource.NamespaceController{},"get:NamespaceListApi"),
		beego.NSRouter("/pod_list", &resource.PodController{},"get:PodListApi"),
        beego.NSRouter("/node_list", &resource.NodeListController{},"get:NodeListApi"),
	)
	beego.AddNamespace(ns)
}
