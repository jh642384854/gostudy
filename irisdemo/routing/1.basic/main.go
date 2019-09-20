package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8090"))
}

func newApp() *iris.Application {
	app := iris.New()
	reigsterErrorCode(app)
	// http://localhost:8090/
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Home index" + ctx.Path())
	})

	// http://localhost:8090/home
	app.Get("/home", func(ctx iris.Context) {
		ctx.HTML("这个方法和app.Handle()方法实现功能一样")
	})

	//一个请求里面有多个handler
	//http://localhost:8090/handler
	app.Get("/handler",doHandler,doFinishHandler)


	//给请求的URL里面绑定参数，并设定参数格式
	// http://localhost:8090/api/article/15
	app.Get("/api/article/{aid:uint64 min(1)}", func(ctx iris.Context) {
		aid,err := ctx.Params().GetUint64("aid");
		if err != nil{
			ctx.Writef("error while trying to parse userid parameter," +
				"this will never happen if :uint64 is being used because if it's not a valid uint64 it will fire Not Found automatically.")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}

		ctx.JSON(map[string]interface{}{
			"article_id":aid,
		})
	})

	//路由分组(定义方式1)
	adminRoutes := app.Party("/admin",adminMiddleware)
	// Done()方法，只针对/admin这个路径生效
	adminRoutes.Done(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("response send to :"+ctx.Path())
	})
	// http://localhost:8090/admin
	adminRoutes.Get("/", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.HTML("Hello form admin")
		ctx.Next()
	})
	// http://localhost:8090/admin/login
	adminRoutes.Get("/login", func(ctx iris.Context) {
		ctx.HTML("admin login")
	})

	//POST http://localhost:8090/admin/login
	adminRoutes.Post("/login", func(ctx iris.Context) {
		ctx.HTML("do admin login")
	})




	//路由分组定义方式2
	api := app.Party("api")
	{
		// http://localhost:8090/api/
		api.Get("/", func(ctx iris.Context) {
			ctx.HTML("API version 1.0")
		})
		//路由分组还可以进行嵌套(无限嵌套)
		useApi := api.Party("/users")
		{
			// http://localhost:8090/api/users
			useApi.Get("/", func(ctx iris.Context) {
				ctx.Writef("all users")
			})
			// http://localhost:8090/api/users/15
			useApi.Get("/{userid:int}", func(ctx iris.Context) {
				ctx.Writef("user with id :%s",ctx.Params().Get("userid"))
			})
		}
	}

	//通配符操作
	wildcardSubdomain := app.Party("*.")
	{
		wildcardSubdomain.Get("/", func(ctx iris.Context) {
			ctx.Writef("Subdomain can be anything, now you're here from: %s", ctx.Subdomain())
		})
	}
	return app
}

func adminMiddleware(ctx iris.Context)  {
	ctx.Next()
}

func doHandler(ctx iris.Context)  {
	ctx.Writef("就像内联处理程序一样，但是它可以被项目中的任何地方的其他包使用")
	ctx.Values().Set("url","https://github.com/kataras/iris#-people")
	ctx.Next()
}

func doFinishHandler(ctx iris.Context)  {
	url := ctx.Values().GetString("url")
	ctx.Application().Logger().Infof("url:"+url)
	ctx.Writef("do send")
}

//应用错误处理
func reigsterErrorCode(app *iris.Application)  {
	app.OnErrorCode(iris.StatusNotFound,notFoundHandeler)
	app.OnErrorCode(iris.StatusInternalServerError,interalServerError)
}

func notFoundHandeler(ctx iris.Context)  {
	ctx.HTML("请求的URL地址没有匹配")
}

func interalServerError(ctx iris.Context)  {
	ctx.HTML("当前服务器内部错误")
}