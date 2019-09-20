package dbops

import (
	"log"
)

func AddVideoCycleRecord(videoid string) error  {
	sqlStr := "INSERT INTO jh_video_recycle (video_id) VALUES(?)"
	stmtInsert,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Printf("Prepare sql error:%v",err)
		return err
	}
	defer stmtInsert.Close()
	_,err = stmtInsert.Exec(videoid)
	if err != nil{
		log.Printf("Exec sql error:%v",err)
		return err
	}
	return nil
}