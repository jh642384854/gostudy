package taskrunner

import (
	"dev/video/scheduler/dbops"
	"errors"
	"log"
	"os"
	"sync"
)

//这个里面就是定义实现dispatch和execute的实现逻辑，不过这里的逻辑需要和数据库进行交互

//分发任务
func VideoClearDispatcher(dc dataChan) error {
	rows,err := dbops.GetCycleVideos(10)
	if err != nil{
		return err
	}
	if len(rows) == 0{
		return errors.New("no rows")
	}
	for _, value := range rows {
		dc <- value
	}
	return nil
}

//执行任务
func VideoClearExecutor(dc dataChan)  error {
	//用来记录错误
	errMap := &sync.Map{}
	//定义错误
	var err error
	forloop:
	for  {
		select {
		case videoID := <- dc:
			go func(id interface{}) {
				//删除视频实体文件
				err := deleteVideoFile(id.(string))
				if err != nil{
					errMap.Store(id,err)
					return
				}
				//删除数据记录
				err = dbops.DeleteCycleVideos(videoID.(string))
				if err != nil{
					errMap.Store(id,err)
					return
				}
			}(videoID)
		default:
			break forloop
		}
	}
	//取得记录的错误信息
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil{
			return false
		}
		return true
	})
	return err
}

//删除视频的实体文件
func deleteVideoFile(videoID string) error {
	err := os.Remove(VIDEO_PATH+videoID)
	if err != nil{
		log.Printf("Delete Video File fail:%v",videoID)
		return err
	}
	return nil
}