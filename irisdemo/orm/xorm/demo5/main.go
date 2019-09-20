package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"

	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

/**
	缓存使用
 */

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口

func main() {
/*
	//下面代码只是测试使用
	var user Users
	fmt.Println(&user,new(Users))  //这两个是相等的
*/
	dbenginer := InitDb()

	defer dbenginer.Close()

	//先获取数据
	var user Users
	hasExits,err := dbenginer.ID(5).Get(&user)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(user,hasExits)
	//更新数据
	upUser := new(Users)
	upUser.UserName = "LiSi_iris_new"
	upUser.Version = 2
	//res,err := dbenginer.Table(new(Users)).ID(5).Update(map[string]interface{}{"user_name":"LiSi_iris","version":3})   //这个操作会失败，是反射引起的错误，要好好复习下反射的内容
	res,err := dbenginer.ID(5).Update(upUser)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(res)
	//再次获取数据，看是否是缓存数据，看更新操作后，是否会更新缓存内容？这个答案是会更新缓存内容
	var user2 Users
	hasExits,err = dbenginer.ID(5).Get(&user2)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(user2)
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

	//缓存设置
	cacher := xorm.NewLRUCacher2(xorm.NewMemoryStore(),7200*time.Second,1000)
	dbenginer.SetDefaultCacher(cacher)

	//开启部分表的缓存
	dbenginer.MapCacher(new(Users),cacher)

	//禁用部分表的缓存
	dbenginer.MapCacher(new(Articles),nil)


	return dbenginer
}

type Users struct {
	Uid int `xorm:"int notnull pk autoincr"`  //这里故意将主键的字段id不是设置为id,而是uid，就是为了测试ID()方法
	UserName string `xorm:"varchar(50)"`
	CreatedTime int `xorm:"created"`  //在插入数据的时候，会自动插入到数据表中
	UpdatedTime int `xorm:"updated"`  //在Insert(), InsertOne(), Update()方法被调用时，updated标记的字段将会被自动更新为当前时间
	DeletedTime time.Time `xorm:"deleted"`  //要实现软删除的功能，这个字段类型必须是time.Time，而对应到mysql的字段类是datetime类型
	Version int `xorm:"version"`
}

//定义一个模型类
type Articles struct {
	Id int `xorm:"int notnull pk autoincr"`
	Title string `xorm:"varchar(255)"`
	Author string `xorm:"varchar(255)"`
	Status int `xorm:"int"`
	Views int `xorm:"int"`
}