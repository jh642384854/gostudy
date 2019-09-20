package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
)

func main() {

	app := iris.New()
	// app.Use() 全局使用中间件
	app.Use(recover2.New())
	app.Use(logger.New())

	app.Handle("GET","/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome Iris</h1>")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message":"Hello Iris!"})
	})


	app.Run(iris.Addr(":8080"),iris.WithoutServerError(iris.ErrServerClosed))
}