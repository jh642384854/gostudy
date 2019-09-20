package utils

import (
	"dev/video/scheduler/dbops"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func VideoDeleteCycleHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	videoid := p.ByName("video-id")

	if len(videoid) == 0{
		SendMessage(w,400,"video id should not be empty")
		return
	}

	err := dbops.AddVideoCycleRecord(videoid)
	if err != nil{
		SendMessage(w,500,"Internal Server error")
		return
	}
	SendMessage(w,200,"successFully")
	return
}