package main

import (
	"github.com/kataras/iris"
)

/**
	这个例子目前还有些疑问

 */

func main() {

	app := iris.New()

	apiv1 := app.Party("/api/v1")
	// 只有在myFilter1为true的情况下，后面的handler才会被执行，如果只有一个handler，那这个handler会被执行，如果后面有多个handler，则会跳跃一个，从二个handler开始执行，不明白为什么会这样来设计？
	myMiddleware := iris.NewConditionalHandler(myFilter1,myHandler1)

	apiv1Router := apiv1.Party("/users",myMiddleware)
	// http://localhost:8080/api/v1/users
	// http://localhost:8080/api/v1/users?admin=true
	apiv1Router.Get("/", func(ctx iris.Context) {
		ctx.HTML("requestd:<b>/api/v1/users</b>")
	})

	app.Run(iris.Addr(":8090"))
}


func myFilter1(ctx iris.Context) bool  {
	ok,_ := ctx.URLParamBool("admin")
	return ok
}

func myHandler1(ctx iris.Context)  {
	//ctx.Application().Logger().Infof("admin: %s", ctx.Params())
	ctx.HTML("<h1>Hello Admin 11111</h1>")
	ctx.Next()
}

func myHandler2(ctx iris.Context)  {
	ctx.HTML("<h1>Hello Admin</h1>")
	ctx.Next()
}

func myHandler3(ctx iris.Context)  {
	ctx.HTML("<h1>myHandler3</h1>")
	ctx.Next()
}
