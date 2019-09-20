package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
)

func main() {

	app := newApp()
	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {

	app := iris.New()
	app.Logger().Info()
	app.Logger().SetLevel("debug")

	//mvc.Configure()函数接受router.Party变量和以mvc.Application作为形参的方法作为参数
	mvc.Configure(app.Party("/basic"),basicMVC)

	return app
}

func basicMVC(mvcApp *mvc.Application)  {
	//给mvc路由注册中间件
	mvcApp.Router.Use(func(ctx context.Context) {
		ctx.Application().Logger().Infof("Path:%s",ctx.Path())
		ctx.Next()
	})

	//给MVC注册一些依赖
	mvcApp.Register(
		sessions.New(sessions.Config{}).Start,
		&preFixedLogger{ prefix:"DEV"},
	)
	mvcApp.Handle(new(basicController))

	//在mvc里面进行分组
	mvcApp.Party("/sub").Handle(new(basicsubController))
}

type LoggerService interface {
	Log(string)
}

type preFixedLogger struct {
	prefix string
}
//实现了LoggerService定义的方法，也就是说preFixedLogger这个实现了LoggerService这个接口
func (p *preFixedLogger) Log(msg string)  {
	fmt.Printf("%s:%s\n",p.prefix,msg)
}

//定义一个Controller
type basicController struct {
	Logger LoggerService
	Session *sessions.Session
}
// http://localhost:8080/basic/custom
// http://localhost:8080/basic/custom2
// mvc.BeforeActivation它在控制器的依赖项绑定到字段或输入参数之前调用，但在服务器运行之前调用。它被用来自定义一个控制器如果需要在控制器内部，它被每个应用程序调用一次。
func (c *basicController) BeforeActivation(b mvc.BeforeActivation)  {
	//HandleMany类似于“Handle”，但可以在同一个控制器的方法上注册多个由空格分隔的路径和HTTP方法路由。
	// 请注意，如果控制器的方法输入参数是路径参数依赖项，那么它们应该与每个给定路径匹配。
	b.HandleMany("GET","/custom /custom2","Custom")
}

//它在“激活前”之后调用，在控制器的依赖项绑定到字段或输入参数之后调用，但在服务器运行之前调用。
func (c *basicController) AfterActivation(b mvc.AfterActivation)  {
	if b.Singleton() {
		panic("basicController 必须是无状态的，在一个请求范围内，我们有一个依赖于上下文的“会话”")
	}
}
// http://localhost:8080/basic
func (c *basicController) Get() string  {
	count := c.Session.Increment("count",1)
	body := fmt.Sprintf("您是当前访问的[:%d]位访问者",count)
	c.Logger.Log(body)
	return body
}

func (c *basicController) Custom() string {
	return "custom"
}

//在MVC中的应用分组
type basicsubController struct {
	Session *sessions.Session
}

func (s *basicsubController) Get() string  {
	count := s.Session.GetIntDefault("count",1)
	return fmt.Sprintf("basicsubController,count:%d \n",count)
}