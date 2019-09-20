package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

/**
	SQL事件
 */

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口

func main() {

	dbenginer = InitDb()
	defer dbenginer.Close()

	article := new(Articles)
	article.Status = 1
	article.Author = "go_iris"
	article.Views = 1
	article.Title = "why are you study iris?"

	res,err := dbenginer.InsertOne(article)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(res)

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

	//定义表的映射关系金和表前缀
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{},"jh_")
	dbenginer.SetTableMapper(tableMapper)

	return dbenginer
}


//定义一个模型类
type Articles struct {
	Id int `xorm:"int notnull pk autoincr"`
	Title string `xorm:"varchar(255)"`
	Author string `xorm:"varchar(255)"`
	Status int `xorm:"int"`
	Views int `xorm:"int"`
}

/**
	在Articles结构体执行插入之前进行的操作
 */
func (a *Articles) BeforeInsert()  {
	fmt.Println("Articles BeforeInsert")
}
/**
	在Articles结构体执行插入之后进行的操作
 */
func (a *Articles) AfterInsert()  {
	fmt.Println("Articles AfterInsert")
}