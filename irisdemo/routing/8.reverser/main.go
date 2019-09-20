package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

func main() {

	app := iris.New()

	rv := router.NewRoutePathReverser(app)

	myroute := app.Get("/anything/{anthingparamter:path}", func(ctx context.Context) {
		paramValue := ctx.Params().Get("anthingparamter")
		ctx.Writef("The path after /anthing is :%s",paramValue)
	})

	myroute.Name = "myroute"

	app.Get("/reverse_myroute", func(ctx context.Context) {
		myrouteRequestPath := rv.Path(myroute.Name,"any/path")
		ctx.HTML("Should be <b>/anything/any/path</b>"+myrouteRequestPath)
	})

	app.Get("/execute_myroute", func(ctx context.Context) {
		ctx.Exec("GET","/anything/any/path")
	})



	app.Run(iris.Addr(":8080"))

}
