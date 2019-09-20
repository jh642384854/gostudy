package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func newApp() *iris.Application  {
	app := iris.New()
	//为应用设置favicon图标
	app.Favicon("./dist/favicon.ico")
	//app.HandleDir()函数用来处理访问地址和实际项目目录的映射关系
	//iris.DirOptions的配置在src/github.com/kataras/iris/core/router/fs.go:24
	app.HandleDir("/","./dist",iris.DirOptions{
		IndexName:"/index.html",
		Gzip:true,
		ShowList:false,
	})

	//下载文件
	app.Get("/download", func(ctx context.Context) {
		file := "./dist/files/first.zip"
		ctx.SendFile(file,"irisdemo.zip")
	})

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}