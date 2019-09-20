package main

import (
	"dev/video/streamserver/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
	这个组件主要用来实现视频的播放和视频的上传
 */

//创建一个流控过滤器的中间件
type LimiterMiddlerWare struct {
	router *httprouter.Router
	ConnLimiter *utils.ConnLimiter
}

func NewLimiterMiddlerWare(count int,r *httprouter.Router) * LimiterMiddlerWare  {
	return &LimiterMiddlerWare{
		router:r,
		ConnLimiter:utils.NewConnLimiter(count),
	}
}

func (limit *LimiterMiddlerWare) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	if ! limit.ConnLimiter.GetConn() {
		utils.SendErrorResponse(w,http.StatusTooManyRequests,"Too Many Requests")
		return
	}
	limit.router.ServeHTTP(w,r)
	defer limit.ConnLimiter.ReleaseConn()
}

//路由注册
func RetisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:video-id",utils.StreamHandler)
	router.POST("/upload/:video-id",utils.UploadHandler)
	return router
}

func main() {
	r := RetisterHandler()
	mr := NewLimiterMiddlerWare(2,r)
	http.ListenAndServe(":9000",mr)
}
