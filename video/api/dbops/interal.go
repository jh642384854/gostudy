package dbops

import (
	"dev/video/api/defs"
	"strconv"
	"sync"
)

/**
sessions操作
 */
//插入
func InsertSession(TTL int64,session_id,login_name string) error {
	sqlStr := "INSERT INTO jh_sessions (session_id,TTL,login_name) VALUES(?,?,?)"
	stmtInsert,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return err
	}
	_,err = stmtInsert.Exec(session_id,TTL,login_name)
	if err != nil{
		return err
	}
	defer stmtInsert.Close()
	return nil
}
//获取单个session信息
func RetrieveSession(session_id string) (*defs.Session,error) {
	var session = &defs.Session{}
	sqlStr := "SELECT TTL,login_name FROM jh_sessions WHERE session_id = ?"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	defer stmtSelect.Close()
	row := stmtSelect.QueryRow(session_id)
	var TTL,login_name string
	err = row.Scan(&TTL,&login_name)
	if err != nil{
		return nil,err
	}
	//将TTL转换为int64
	if ttlRes,err := strconv.ParseInt(TTL,10,64); err != nil{
		return nil,err
	}else{
		session.TTL = ttlRes
		session.LoginName = login_name
		session.SessionID = session_id
	}
	return session,nil
}
//获取所有session信息
func RetrieveAllSession() (*sync.Map,error) {
	m := &sync.Map{}
	sqlStr := "SELECT * FROM jh_sessions"
	stmtSelect,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return nil,err
	}
	defer stmtSelect.Close()
	rows,err := stmtSelect.Query()
	if err != nil{
		return nil,err
	}
	for rows.Next() {
		var session_id,TTL,login_name string
		if err := rows.Scan(&session_id,&TTL,&login_name); err != nil{
			break
		}
		if ttlres,err := strconv.ParseInt(TTL,10,64);err != nil{
			session := &defs.Session{
				SessionID:session_id,
				TTL:ttlres,
				LoginName:login_name,
			}
			m.Store(session_id,session)
		}
	}
	return m,nil
}
//删除session信息
func DeleteSession(session_id string) error  {
	sqlStr := "DELETE FROM jh_sessions WHERE  session_id = ?"
	stmtDelete,err := dbConn.Prepare(sqlStr)
	if err != nil{
		return err
	}
	defer stmtDelete.Close()
	_,err = stmtDelete.Exec(session_id)
	if err != nil{
		return err
	}
	return nil
}