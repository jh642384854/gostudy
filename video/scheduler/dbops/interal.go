package dbops

import "log"

/**
	获得被删除记录的视频记录信息
 */
func GetCycleVideos(count int) ([]string,error) {
	var ids []string
	sqlStr := "SELECT video_id FROM jh_video_recycle LIMIT ?"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Printf("sql : Prepare sql error")
		return ids,err
	}
	defer stmtSelect.Close()
	result,err := stmtSelect.Query(count)
	if err != nil{
		log.Printf("sql : Exec sql error")
		return ids,err
	}

	for result.Next() {
		var videoid string
		if err := result.Scan(&videoid); err != nil{
			log.Printf("sql : Scan sql error")
			return ids,err
		}
		ids = append(ids,videoid)
	}
	return ids,nil
}

//删除一条要被删除的视频记录信息
func DeleteCycleVideos(video_id string) error {
	sqlStr := "DELETE FROM jh_video_recycle WHERE video_id = ?"
	stmtDel,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Printf("sql : Prepare sql error")
		return err
	}
	defer stmtDel.Close()
	_,err = stmtDel.Exec(video_id)
	if err != nil{
		log.Printf("sql : Exec sql error")
		return err
	}
	return nil
}