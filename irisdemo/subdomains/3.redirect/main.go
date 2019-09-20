package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

const ADDR = "jhgo.com:80"

func main() {
	app := newApp()
	app.Run(iris.Addr(ADDR))
}

func newApp() *iris.Application  {
	
	app := iris.New()
	app.Get("/", func(ctx context.Context) {
		ctx.Writef("这个请求永远不会被执行到")
	})

	www := app.Subdomain("www")
	www.Get("/",index)


	www.PartyFunc("/users", func(p router.Party) {
		p.Get("/",userIndex)
		p.Get("/login",userLogin)
	})
	//这个会将jhgo.com通过301重定向到www.jhgo.com
	app.SubdomainRedirect(app,www)

	return app
}

func index(ctx iris.Context)  {
	ctx.Writef("这是www.jhgo.com首页显示内容")
}

func userIndex(ctx iris.Context)  {
	ctx.Writef("这是www.jhgo.com/user页面显示内容")
}

func userLogin(ctx iris.Context)  {
	ctx.Writef("这是www.jhgo.com/users/login页面显示内容")
}
