package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

/**
	这里也是使用递归的方式来实现，不过这里只是查询一次数据库
 */

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口


func main() {
	category := make([]Category,0)
	dbenginer = InitDb()
	if err := dbenginer.OrderBy("sort").Find(&category); err != nil{
		fmt.Println(err.Error())
		return
	}
	category2 := new(Category)
	categories := category2.CategoryArrNew(category,0,0)
	categoriesJson,err := json.Marshal(categories)
	if err != nil{
		return
	}
	fmt.Println(string(categoriesJson))
}

/**
	参考文档：http://note.youdao.com/noteshare?id=94e30c5323cd7fba6b0dec26fb03c779&sub=2E5D03949EAE42B08AF9B8FAB44E6A7F这个来实现的
 */
func (cate *Category) CategoryArrNew(category []Category,pid ,level int) []Category  {
	list := []Category{}
	for _, value := range category {
		if value.Pid == pid{
			value.Level = level
			value.Son = cate.CategoryArrNew(category,value.Id,level+1)
			list = append(list,value)
		}
	}
	return list
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

type Category struct {
	Id int `json:"id" xorm:"id"`
	Pid int `json:"pid" xorm:"pid"`
	Name string `json:"name" xorm:"name"`
	Sort int `json:"sort" xorm:"sort"`
	Level int `json:"level" xorm:"-"`
	Son []Category `json:"son" xorm:"-"`
}

