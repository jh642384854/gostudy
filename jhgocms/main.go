package main

import (
	"fmt"
	"github.com/betacraft/yaag/irisyaag"
	"github.com/betacraft/yaag/yaag"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"jhgocms/controller"
	"jhgocms/datasource"
)

func main() {

	app := initApp()
	
	//①、系统配置
	configApp(app)

	//②、APIDOC处理
	//apiDocConfig(app)

	//③、MVC绑定
	mvcHandler(app)

	app.Run(iris.Addr(":8080"))
}
/**
	系统初始化配置
 */
func initApp() *iris.Application {
	app := iris.New()

	//设置日志级别
	app.Logger().SetLevel("debug")

	//静态资源设置
	app.Favicon("./public/system/favicon.ico")
	app.HandleDir("/","./public/system",iris.DirOptions{
		IndexName:"/index.html",
		Gzip:true,
		ShowList:false,
	})

	//注册视图文件
	app.RegisterView(iris.HTML("./public/system",".html"))
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	return app
}

/**
	系统应用配置
 */
func configApp(app *iris.Application)  {

}

/**
	yaag doc api注册处理
 */
func apiDocConfig(app *iris.Application)  {
	yaag.Init(&yaag.Config{
		On:true,
		DocTitle: "Iris",
		DocPath:  "apidoc.html",
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})
	app.Use(irisyaag.New())
}
/**
	mvc初始化设置
*/
func mvcHandler(app *iris.Application)  {


	mysqlDbEngine := datasource.NewMysqlEngine()
	fmt.Println(mysqlDbEngine)

	//user路由组
	adminUserRoutes := mvc.New(app.Party("/user",adminMiddleware))
	adminUserRoutes.Handle(new(controller.AdminUserController))
}

func adminMiddleware(ctx iris.Context)  {
	ctx.Next()
}
