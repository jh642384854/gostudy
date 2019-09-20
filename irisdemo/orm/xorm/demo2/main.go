package main

import (
	"fmt"
	"time"
	"xorm.io/core"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口

func main() {
	dbenginer := InitDb()

	defer dbenginer.Close()
	// 查询操作的基本示例
	//FindDemo(dbenginer)
	//InsertDemo(dbenginer)
	//UpdateDemo(dbenginer)
	DeleteDemo(dbenginer)
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

//定义一个模型类
type Articles struct {
	Id int `xorm:"int notnull pk autoincr"`
	Title string `xorm:"varchar(255)"`
	Author string `xorm:"varchar(255)"`
	Status int `xorm:"int"`
	Views int `xorm:"int"`
}

type Users struct {
	Uid int `xorm:"int notnull pk autoincr"`  //这里故意将主键的字段id不是设置为id,而是uid，就是为了测试ID()方法
	UserName string `xorm:"varchar(50)"`
	CreatedTime int `xorm:"created"`  //在插入数据的时候，会自动插入到数据表中
	UpdatedTime int `xorm:"updated"`  //在Insert(), InsertOne(), Update()方法被调用时，updated标记的字段将会被自动更新为当前时间
	DeletedTime time.Time `xorm:"deleted"`  //要实现软删除的功能，这个字段类型必须是time.Time，而对应到mysql的字段类是datetime类型
	Version int `xorm:"version"`
}


/**
	查询的基本操作
 */
func FindDemo(dbenginer *xorm.Engine)  {
	//①、获取单条记录。使用Where()方法来指定其中的查询条件来实现各种条件的查询
	var article Articles
	//方式一：Get()方法，用来获取数据记录，返回两个数据，第一个参数返回该查询记录是否存在，第二个参数返回是否有错误发生
	hasExist,err := dbenginer.Where("id = ?",1).Get(&article);
	if err != nil{
		panic(err.Error())
	}
	if hasExist {
		fmt.Println(article)
	}
	//方式二：使用ID()方法，这个就必须在模型的字段的xorm的tag标签设为主键，即需要有pk的说明
	var article2 Articles
	hasExist,err = dbenginer.ID(2).Get(&article2)
	fmt.Println(hasExist,err)
	fmt.Println(article2)

	//下面的这个例子是用来测试当表结构定义的主键字段并不是id，依然使用ID()方法是否能够查询到数据
	var user Users
	dbenginer.ID(1).Get(&user)
	fmt.Println(user)

	//②、获取多条记录：使用Find()方法，简单粗暴，获取当前操作表的所有记录(不分页)
	articles := make([]Articles,0)
	if err := dbenginer.Find(&articles); err != nil{
		panic(err.Error())
	}
	for _, article := range articles {
		fmt.Println(article)
	}

	//带分页的方式获取数据
	articles2 := make([]Articles,0)
	dbenginer.Limit(5,0).Find(&articles2)
	for _, article := range articles2 {
		fmt.Println(article)
	}
	fmt.Println()
	//Find()带条件的查询
	articles3 := make([]Articles,0)
	dbenginer.Where("status = ?",0).Find(&articles3)
	for _, article := range articles3 {
		fmt.Println(article)
	}
	fmt.Println()

	//或是下面的这种带条件的查询   Find()第二个参数使用例子，只能是借鉴 https://www.2cto.com/kf/201905/810039.html
	articles4 := make([]Articles,0)
	acondiont := new(Articles)
	acondiont.Status = 1
	dbenginer.Find(&articles4,acondiont)
	for _, article := range articles4 {
		fmt.Println(article)
	}
	fmt.Println()

	//③、查询指定字段的数据。使用Cols()方法用来指定要查询的字段信息
	var article3 Articles
	dbenginer.ID(5).Cols("title","author").Get(&article3)
	fmt.Println(article3)

	//④、聚合函数的使用
	//1.计算记录总和
	count,err := dbenginer.Count(new(Articles))
	if err != nil{
		panic(err.Error())
	}
	fmt.Printf("sum:%d \n",count)
	//2.计算字段总和
	sum,err := dbenginer.SumInt(new(Articles),"views")
	fmt.Printf("sum:%d \n",sum)
	//3.xorm并没有实现max、min、avg等等聚会函数的封装


	//5.Rows()方法使用
	articleRow := new(Articles)
	rows,err := dbenginer.Rows(articleRow)
	for rows.Next(){
		if err := rows.Scan(articleRow); err != nil{
			panic(err.Error())
		}
		fmt.Println(*articleRow)
	}


}

/**
	插入操作
 */
func InsertDemo(dbenginer *xorm.Engine)  {

	//1.单条记录插入操作
	article := new(Articles)
	article.Status = 1
	article.Author = "dawang"
	article.Title = "xorm study"
	article.Views = 0
	//InsertOne()返回两个值，一个是是否插入成功的一个整数，一个是错误。注意：这个整数并不是返回插入后最后自增主键ID，但是可以使用这个模型的主键字段可以获取最后自增的id
	/*aid,err := dbenginer.InsertOne(article)
	if err != nil{
		panic(err.Error())
	}
	fmt.Printf("lastInsertID :%d，aid:%d",article.Id,aid)*/

	//2.多条记录的插入操作。Insert()方法可以用来插入单条记录、多条记录、多条不同表的记录、多条同表的记录

	articles := make([]*Articles,2)

	articles1 := new(Articles)
	articles1.Title = "go study"
	articles1.Views = 0
	articles1.Author = "ruixue"
	articles1.Status = 1
	articles[0] = articles1

	articles[1] = article
	//插入多条同表的记录
	//dbenginer.Insert(articles)



	users := make([]*Users,2)
	users[0] = new(Users)
	users[0].UserName = "wangmaz"

	users[1] = new(Users)
	users[1].UserName = "lisi"
	//插入多条同表的记录
	//dbenginer.Insert(users)

	//插入多条记录，但是并不是同表
	//dbenginer.Insert(articles,users)


	userCreated := new(Users)
	userCreated.UserName = "xiaoeme"
	dbenginer.InsertOne(userCreated)
	fmt.Println(userCreated)

}

/**
	更新操作
 */
func UpdateDemo(dbenginer *xorm.Engine)  {
	//1、普通的字段更新.下面设定了几个字段，就会更新哪几个字段。
	//①、方式一：接受某个结构体指针的方式
	article := new(Articles)
	article.Author = "lixuh"
	//dbenginer.ID(3).Update(article)
	//②、方式二：接受Map[string]interface{}类型，这种类型，就需要单独通过Table()方法来指定要操作的是哪个数据表
	//dbenginer.Table(new(Users)).ID(3).Update(map[string]string{"user_name":"goubuli"})

	//2、更新字段值为0的操作方式
	//①、方式1：
	article2 := new(Articles)
	article2.Status = 0
	//dbenginer.ID(1).Cols("status").Update(article2)

	//②、方式2：这里必须要使用Table()方法来指定操作的是那个数据表，不然是操作不了的，因为没有关联具体的哪个数据表
	//dbenginer.Table(new(Articles)).ID(2).Update(map[string]interface{}{"status":0})

	//3.乐观锁的使用，在使用乐观锁的时候，必须要传递Version字段。
	user := new(Users)
	user.UserName = "mahaige"
	user.Version = 1
	//dbenginer.InsertOne(user)
	res,err := dbenginer.ID(7).Update(user)
	if err != nil{
		panic(err.Error())
	}
	fmt.Print(res)


}

/**
	删除操作
 */
func DeleteDemo(dbenginer *xorm.Engine)  {
	//1.物理删除
/*
	res,err := dbenginer.ID(12).Delete(new(Articles))
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(res)
*/
	//2.逻辑删除(软删除)
/*
	var user Users
	res,err := dbenginer.ID(4).Delete(user)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(res)
*/
	//在获取被删除的记录，这样就是获取不到了的
	user := new(Users)
	resBool,err := dbenginer.ID(4).Get(user)
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(resBool)

	//获取被标记为软删除的记录.这个需要用到Unscoped()方法
/*
	dbenginer.ID(4).Unscoped().Get(user)
	fmt.Println(user)
*/
	//将软删除的记录进行物理删除。这个也需要用到Unscoped()方法
	dbenginer.ID(4).Unscoped().Delete(user)

}