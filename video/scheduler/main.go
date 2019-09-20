package scheduler

import (
	"dev/video/scheduler/taskrunner"
	"dev/video/scheduler/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
	这个组件的主要功能:
	①：将把要是删除的视频存放到相应的数据表中(类似回收站)
	②：读取视频回收站数据表中的数据，进行分发任务，最后执行任务
 */

func registerHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-recode/:video-id",utils.VideoDeleteCycleHandler)
	return router
}

func main() {
	taskrunner.Start()
	r := registerHandler()
	http.ListenAndServe(":8080",r)
}