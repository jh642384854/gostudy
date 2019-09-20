package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()

	//定义视图
	template := iris.HTML("./templates",".html")
	template.Reload(true)

	template.AddFunc("green", func(s string) string {
		return "Greetings "+ s + "!"
	})

	//为应用注册模版应用
	app.RegisterView(template)
	app.Get("/", func(ctx context.Context) {
		ctx.Gzip(true)
		ctx.ViewData("Title","iris title")
		ctx.ViewData("Name","iris name")
		ctx.View("index.html")
	})

	return app
}
