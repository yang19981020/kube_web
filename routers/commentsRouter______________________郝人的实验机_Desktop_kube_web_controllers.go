package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["kube_web/controllers:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers:UserController"],
        beego.ControllerComments{
            Method: "Pass",
            Router: "/updatePass",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
