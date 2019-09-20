package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

func main() {
	app := newApp()
	app.Run(iris.Addr("jhgo.com:80"))
}

/**
	接收以下URL地址请求：

	http://jhgo.com
	http://jhgo.com/about
	http://ijhgo.com/contact
	http://jhgo.com/api/users
	http://jhgo.com/api/users/42

	http://www.jhgo.com
	http://www.jhgo.com/hi
	http://www.jhgo.com/about
	http://www.jhgo.com/contact
	http://www.jhgo.com/api/users
	http://www.jhgo.com/api/users/42
 */

func newApp() *iris.Application  {
	app := iris.New()

	app.Get("/",info)
	app.Get("/about",info)
	app.Get("/contact",info)

	app.PartyFunc("/api/users", func(p router.Party) {
		p.Get("/",info)
		p.Get("/{id:uint64}",info)
	})

	//如果不做这个处理，使用www来进行访问都会出问题。
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			www.Handle(r.Method,r.Tmpl().Src,r.Handlers...)
		}
		www.Get("/hi", func(ctx context.Context) {
			ctx.Writef("Hi form www.jhgo.com")
		})
	}

	return app
}

func info(ctx iris.Context)  {
	method := ctx.Method()
	path := ctx.Path()
	subdomain := ctx.Subdomain()//注意：如果iris.Addr("www.jhgo.com:80")是这样绑定的，那通过ctx.Subdomain()获取的内容就是空的

	ctx.Writef("\n info \n\n")
	ctx.Writef("Method :%s，Path:%s，subdomain:%s",method,path,subdomain)
}