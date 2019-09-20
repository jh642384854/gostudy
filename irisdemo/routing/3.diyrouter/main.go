package main

import (
	"strings"


	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
)

/**
	自定义的router
 */

func main() {
	app := iris.New()

	app.Get("/hello/{name}", func(ctx context.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("Hello %s \n",name)
	})

	app.Get("/cs/{num:uint64 min(10) else 400}", func(ctx context.Context) {
		num := ctx.Params().GetUint64Default("num",0)
		ctx.Writef("num is :%d\n",num)
	})

	myCustomerRouter := new(customRouter)
	app.BuildRouter(app.ContextPool,myCustomerRouter,app.APIBuilder,true)

	app.Run(iris.Addr(":8090"))
}

//这个类实现了RequestHandler这个接口的所有方法
type customRouter struct {
	provider router.RoutesProvider
}

//以下三个方法都是src/github.com/kataras/iris/core/router/handler.go:17中定义的RequestHandler这个接口的实现
func (r *customRouter) HandleRequest(ctx context.Context)  {
	path := ctx.Path()
	ctx.Application().Logger().Infof("Requested resource path: %s", path)

	parts := strings.Split(path,"/")[1:]
	staticPath := "/"+parts[0]
	for _, route := range r.provider.GetRoutes() {
		if strings.HasPrefix(route.Path,staticPath) && route.Method == ctx.Method(){
			paramParts := parts[1:]
			for _, paramValue := range paramParts {
				for _, p := range route.Tmpl().Params {
					ctx.Params().Set(p.Name,paramValue)
				}
			}
			ctx.SetCurrentRouteName(route.Name)
			ctx.Do(route.Handlers)
			return
		}
	}
}

func (r *customRouter) Build(provider router.RoutesProvider) error  {
	for _,route := range provider.GetRoutes(){
		route.BuildHandlers()
	}
	r.provider = provider
	return nil
}

func (r *customRouter) RouteExists(ctx context.Context,method,path string) bool  {
	return false
}