package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beegoproject/controllers:AnnotationController"] = append(beego.GlobalControllerRouter["beegoproject/controllers:AnnotationController"],
        beego.ControllerComments{
            Method: "Demo1",
            Router: `/annotation/demo1`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
