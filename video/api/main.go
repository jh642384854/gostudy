package main


import (
	"dev/video/api/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//定义一个中间件，这个中间件必须要实现src/net/http/server.go中Handler接口的ServeHTTP()方法。
type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	//合法性校验
	utils.ValidateUserSession(r)
	m.r.ServeHTTP(w,r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler  {
	m := middleWareHandler{}
	m.r = r
	return m
}



func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()
	router.POST("/user",utils.CreateUser)
	router.POST("/login/:user_name",utils.Login)
	return router
}

func main() {
	r := RegisterHandlers()
	//注册中间件拦截器
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000",mh)
}

/**
	处理流程如下：
	main -> middleware ->defs(message,error) -> handlers -> dbops -> response
	main：代表主入口
	middleware：代表中间件，用来对请求进行合法性和权限校验
	defs(message,error)：这个是表用来定义消息
	handlers：这个就是对应请求处理逻辑
	dbops：请求处理逻辑可能会涉及到到的数据库相关操作
	response：将请求处理的结果进行返回。
 */
