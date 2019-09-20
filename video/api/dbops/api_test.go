package dbops

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"testing"
)

/**
	init(dblogin,truncate tables)->run
	tests->clear data(truncate tables)
 */

func clearTables()  {
	dbConn.Exec("TRUNCATE jh_users")
	dbConn.Exec("TRUNCATE jh_video_info")
	dbConn.Exec("TRUNCATE jh_comments")
	dbConn.Exec("TRUNCATE jh_sessions")
}

/**
	使用TestMain作为初始化test，并且使用m.Run()来调用其他tests可以完成一些需要初始化操作的testing，比如数据库连接、文件打开、REST服务登录等
	如果没有在TestMain中调用m.Run()方法，则除了TestMain以外的其他tests方法都不会被执行。
 */
func TestMain(m *testing.M)  {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T)  {
	t.Run("Add",TestAddUserCredential)
	t.Run("Get",TestGetUserCredential)
	t.Run("Del",TestDeleteUser)
	t.Run("ReGet",TestGetUserCredential2)
}


func TestAddUserCredential(t *testing.T) {
	t.SkipNow()
	err := AddUserCredential("zhangsan","123465")
	if err != nil{
		t.Errorf("Error of AddUserCredential:%v",err)
	}
}

func TestGetUserCredential(t *testing.T) {
	pwd,err := GetUserCredential("zhangsan")
	if pwd != "123465" || err != nil{
		t.Errorf("Error of GetUserCredential:%v",err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("zhangsan","123465")
	if err != nil{
		t.Errorf("Error of TestDeleteUser:%v",err)
	}
}

func TestGetUserCredential2(t *testing.T) {
	pwd,err := GetUserCredential("zhangsan")
	if err != nil{
		t.Errorf("Error of TestGetUserCredential2:%v",err)
	}
	if pwd != ""{
		t.Errorf("Record not delete")
	}
}

func TestAddVideo(t *testing.T) {
	videoinfo,err := AddVideo(1,"spl")
	if err != nil{
		log.Error(err)
	}
	fmt.Println(videoinfo)
}