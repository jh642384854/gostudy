package main

import (
	"github.com/kataras/iris"
)

func newApp() *iris.Application  {

	app := iris.New()

	// 处理请求方式1
	app.Handle("GET","/getRequest", func(ctx iris.Context) {
		ctx.WriteString("getRequest demo1")
	})

	// 处理请求方式2
	app.Get("/getRequest2", func(ctx iris.Context) {
		ctx.WriteString("getRequest demo2")
	})

	// 获取请求的URL参数(逐个获取URL的请求参数)
	app.Get("/getRequest3", func(ctx iris.Context) {

		param1 := ctx.Params().Get("param1") //这个获取不到值，这是为什么呢？
		param2 := ctx.URLParam("param2")

		app.Logger().Info(param1,param2)
	})

	// 获取请求的URL参数(批量获取url的请求各个参数)
	app.Get("/getQuery", func(ctx iris.Context) {
		var query Query
		if err := ctx.ReadQuery(&query); err != nil{
			panic(err.Error())
		}
		app.Logger().Info("query:",query)
	})

	// 获取post提交的数据1(依次获取表单提交的部分字段值)
	app.Post("/postRequest", func(ctx iris.Context) {

		username := ctx.PostValue("username")
		passwd := ctx.FormValue("passwd")
		app.Logger().Info("username:",username," passwd:",passwd)
	})

	// 获取post提交的数据2(批量获取表单提交的全部数据)
	app.Post("/postForm", func(ctx iris.Context) {
		var article Article
		if err := ctx.ReadForm(&article); err != nil{
			panic(err.Error())
		}
		app.Logger().Info("Article:",article)
	})


	// 获取JSON格式的数据(就是直接传递的是json格式的数据)
	// 测试数据：{"username":"wangwu","passwd":"147852"}
	app.Post("/postJson", func(ctx iris.Context) {

		var user User
		if err := ctx.ReadJSON(&user); err != nil{
			panic(err.Error())
		}
		app.Logger().Info("User:",user)

	})

	// 获取xml格式提交的数据,转换的结构体的映射必须也是xml，而并不是json
	/**
		测试数据格式如下：
		<person>
			<username>wangwu</username>
			<passwd>147852</passwd>
		</person>
	 */
	app.Post("/postXml", func(ctx iris.Context) {

		var person Person
		if err := ctx.ReadXML(&person); err != nil{
			panic(err.Error())
		}
		app.Logger().Info("User:",person)

	})


	return app
}

//需要进行json对象转换的结构体
type User struct {
	Username string `json:"username"`
	Passwd string `json:"passwd"`
}

//需要进行xml对象转换的结构体
type Person struct {
	Username string `xml:"username"`
	Passwd string `xml:"passwd"`
}

//form对象转换
type Article struct {
	Title string
	Author string `form:"user"`
}

type Query struct {
	Uid int `url:"uid"`
	Aid int `url:"aid"`
}

func main() {

	app := newApp()
	app.Run(iris.Addr(":8080"))

}
