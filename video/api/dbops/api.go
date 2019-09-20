package dbops

import (
	"database/sql"
	"dev/video/api/defs"
	"github.com/satori/go.uuid"
	"log"
	time2 "time"
)

/**
	user 数据操作
 */
func AddUserCredential(loginName ,pwd string) error  {
	sqlStr := "INSERT INTO jh_users (login_name,pwd) VALUES(?,?)"
	stmtInsert,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Println("insert sql error:",err.Error())
		return err
	}
	_,err = stmtInsert.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtInsert.Close()
	return nil
}

func GetUserCredential(loginName string) (string,error)  {
	sqlStr := "SELECT pwd FROM jh_users WHERE login_name = ?"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Println("select sql error:",err.Error())
		return "",err
	}
	var pwd string
	err = stmtSelect.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "",err
	}
	defer stmtSelect.Close()
	return pwd,nil
}

func DeleteUser(loginName ,pwd string) error {
	sqlStr := "DELETE FROM jh_users WHERE login_name = ? AND pwd = ?"
	stmtDelete,err := dbConn.Prepare(sqlStr)
	if err != nil{
		log.Println("delete sql error:",err.Error())
		return err
	}
	_,err = stmtDelete.Exec(loginName,pwd)
	if err != nil{
		return err
	}
	defer stmtDelete.Close()
	return nil
}

/**
	video 数据操作
 */
func AddVideo(authid int,name string) (*defs.VideoInfo,error)  {
	uuidv4,_ := uuid.NewV4()
	videoID := uuidv4.String()

	time := time2.Now()
	displayCtime := time.Format("2006-01-02 15:04:05")

	sqlStr := "INSERT INTO jh_video_info (id,author_id,name,display_ctime) VALUES (?,?,?,?)"
	stmtInsert,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	_,err = stmtInsert.Exec(videoID,authid,name,displayCtime)
	if err != nil{
		return nil,err
	}
	defer stmtInsert.Close()

	videoInfo := &defs.VideoInfo{
		Id:videoID,
		AuthorID:authid,
		Name:name,
		DisplayCtime:displayCtime,
	}
	return videoInfo,nil
}

func GetVideo(vid string) (*defs.VideoInfo,error) {
	sqlStr := "SELECT id,author_id,name,display_ctime FROM jh_video_info WHERE id = ?"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	defer stmtSelect.Close()
	row := stmtSelect.QueryRow(vid)
	var videoInfo defs.VideoInfo
	err = row.Scan(&videoInfo.Id,&videoInfo.AuthorID,&videoInfo.Name,&videoInfo.DisplayCtime)
	if err != nil{
		return nil,err
	}
	return &videoInfo,nil
}

func DeleteVideo(vid string) error  {
	sqlStr := "DELETE FROM h_video_info WHERE id = ?"
	stmtDelete,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return err
	}
	defer  stmtDelete.Close()
	_,err = stmtDelete.Exec(vid)
	if err != nil{
		return err
	}
	return nil
}

/**
	comment操作
 */
//添加评论
func AddComment(author_id int,video_id,content string) error {
	//生成唯一的uuid
	uuidv4,_:= uuid.NewV4()
	commentid := uuidv4.String()
	sqlStr := "INSERT INTO jh_comments(id,video_id,author_id,content) VALUES(?,?,?,?)"
	stmtInsert,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return err
	}
	defer stmtInsert.Close()
	_,err = stmtInsert.Exec(commentid,video_id,author_id,content)
	if err != nil{
		return err
	}
	return  nil
}

//获取所有的评论信息
func ListComments(video_id string,from,to int64) ([]*defs.Comment,error) {
	sqlStr := "SELECT c.id,c.video_id,c.content,u.login_name FROM jh_comments AS c INNER JOIN jh_users AS u ON c.author_id = u.id WHERE c.video_id = ? AND c.create_time > FROM_UNIXTIME(?) AND c.create_time <= FROM_UNIXTIME(?)"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	defer stmtSelect.Close()
	records,err := stmtSelect.Query(video_id,from,to)
	if err != nil{
		return  nil,err
	}
	var comments []*defs.Comment
	for records.Next() {
		var id,video_id,content,login_name string
		if err := records.Scan(&id,&video_id,&content,&login_name); err != nil{
			return nil,err
		}
		comment := &defs.Comment{
			ID:id,
			VideoID:video_id,
			Content:content,
			AuthorName:login_name,
		}
		comments = append(comments,comment)
	}
	return comments,nil
}