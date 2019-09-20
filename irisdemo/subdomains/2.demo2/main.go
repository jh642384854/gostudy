package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := newApp()
	app.Run(iris.Addr("jhgo.com:80"))
}

/**
	本例允许访问的URL地址：
	http://admin.jhgo.com
	http://admin.jhgo.com/setting


	http://jhgo.com
	http://jhgo.com/news
 */
func newApp() *iris.Application  {
	app := iris.New()

	app.HandleDir("/images","./dist/static/img/")

	admin := app.Party("admin.")
	{
		admin.Get("/", func(ctx context.Context) {
			ctx.Writef("管理后台首页")
		})
		admin.Get("/setting", func(ctx context.Context) {
			ctx.Writef("管理后台系统设置页面")
		})
	}


	images := app.Party("images.")
	{
		images.Get("/", func(ctx context.Context) {
			ctx.Redirect("http://www.jhgo.com",iris.StatusMovedPermanently)
		})
		//下面的这个语句并没有生效，这是为什么呢？难道app.HandleDir()是全局的？
		app.HandleDir("/images","./dist/static/img/")
	}

	//通配符
	dynamicSubdomains := app.Party("*.")
	{
		dynamicSubdomains.Get("/",dynamicSubdomainHandler)
		dynamicSubdomains.Get("/something",dynamicSubdomainHandler)
		dynamicSubdomains.Get("/something/{param}",dynamicSubdomainHandlerWithParam)
	}

	app.Get("/", func(ctx context.Context) {
		ctx.Writef("网站前台首页")
	})

	app.Get("/news", func(ctx context.Context) {
		ctx.Writef("网站新闻页面")
	})

	return  app
}

func dynamicSubdomainHandler(ctx iris.Context)  {
	subdomain := ctx.Subdomain()
	ctx.Writef("当前请求的二级域名是：%s",subdomain)
}

func dynamicSubdomainHandlerWithParam(ctx iris.Context)  {
	subdomain := ctx.Subdomain()
	ctx.Writef("当前请求的二级域名是：%s，请求的参数是：%s\n",subdomain,ctx.Params().Get("param"))

}