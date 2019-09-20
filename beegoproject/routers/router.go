package routers

import (
	"beegoproject/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.BConfig.RouterCaseSensitive = false

    beego.Router("/", &controllers.MainController{})
    //会根据具体的HTTP请求方式来自动的调用controllers.RouterController这个控制器对应的HTTP的请求方法函数
    //如：如果是get请求的/router这个地址，就会调用controllers.RouterController这个控制器的Get()方法。
    //如果是post请求的/router这个地址，就会调用controllers.RouterController这个控制器的Post()方法。依次类推
    beego.Router("/router",&controllers.RouterController{})

	beego.Router("/router",&controllers.RouterController{},"post:DiyPost")

	beego.Router("/router/diyfun",&controllers.RouterController{},"get:Demo")

    beego.Get("/get", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello get"))
	})

	//加载注解路由
	beego.Include(&controllers.AnnotationController{})
}
