package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"reflect"
	"xorm.io/core"
)

/**
	Query()、QueryString()、QueryInterface()、Exec()方法的使用

 */

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口


func main() {
	dbenginer := InitDb()

	defer dbenginer.Close()

	//执行纯SQL语句
	// Query()、QueryString()、QueryInterface()这三个方法主要用来执行一些查询操作
	sql1 := "select * from jh_articles"
	//①、Query()方法，返回的是[]map[string][]byte这种数据结构，也就是字段数据是二进制格式，如果要查看，还需要进转换，即将[]byte转换为字符串
/*
	res,err := dbenginer.Query(sql1)
	if err != nil{
		panic(err.Error())
	}
	for _, row := range res {
		//*(*string)(unsafe.Pointer(&row["title"][:]))
		fmt.Println(string(row["author"][:]))
	}
*/
	//②、QueryString()方法返回的就是[]map[string]string这种类型，就是直接为字符串，就不需要进行数据格式的转换
/*
	res2,err := dbenginer.QueryString(sql1)
	for _, row := range res2 {
		fmt.Println(row)
	}
*/
	//③、QueryInterface()方法返回的就是[]map[string]interface{}.这个暂时还没有掌握如何将interface{}转换为string
	res3,err := dbenginer.QueryInterface(sql1)
	fmt.Println(err)
	for _, row := range res3 {
		fmt.Println(reflect.TypeOf(row["id"]),reflect.TypeOf(row["author"]))
	}

	//④、Exec()方法：这个方法主要是针对Insert， Update， Delete等操作
	sql2 := "update jh_articles set title = ? where id = ?"
	sqlRes,err := dbenginer.Exec(sql2,"iris title",1)
	fmt.Println(sqlRes.RowsAffected())

}


//创建数据库连接
func InitDb() *xorm.Engine {
	dbenginer,err := xorm.NewEngine("mysql",DATASOURCE)

	if err != nil{
		panic(err.Error())
	}
	if dbenginer.Ping() != nil {
		panic(err.Error())
	}
	//defer dbenginer.Close()

	dbenginer.ShowSQL(true)

	//定义表的映射关系
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{},"jh_")
	dbenginer.SetTableMapper(tableMapper)

	return dbenginer
}