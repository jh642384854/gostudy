package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

/**
	递归实现无限极分类
 */

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口


func main() {
	category := new(Category)
	categories := category.CategoryArr(0)
	categoriesJson,err := json.Marshal(categories)
	if err != nil{
		return
	}
	fmt.Println(string(categoriesJson))

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
	Id int64 `json:"id" xorm:"id"`
	Pid int `json:"pid" xorm:"pid"`
	Name string `json:"name" xorm:"name"`
	Sub []Category `json:"sub" xorm:"-"`
}

func (cate *Category) CategoryArr(pid int64) []Category {
	//category := new([]*Category)
	category := make([]Category,0)
	dbenginer = InitDb()
	if err := dbenginer.Where("pid=?",pid).Find(&category); err != nil{
		fmt.Println(err.Error())
		return nil
	}
	if category == nil{
		return  nil
	}
	categoryNode := []Category{}
	for _, val := range category {
		child := cate.CategoryArr(val.Id)
		node := Category{
			Id:val.Id,
			Name:val.Name,
			Pid:val.Pid,
		}
		node.Sub = child
		categoryNode = append(categoryNode,node)
	}
	return categoryNode
}