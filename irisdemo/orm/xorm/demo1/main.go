package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"os"
)

var dbenginer *xorm.Engine
//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
const DATASOURCE  = "root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8"  //这个可以简写为root:jianghua@/godb?charset=utf8，就是使用默认tcp协议连接本机的3306这个端口
const RDSDATASOURCE = "bjkaifa:VEgGdBaQRzdle7OI@tcp(rm-2zejcj308e99xcbuu.mysql.rds.aliyuncs.com:3306)/linghang?charset=utf8"

func main() {
	dbenginer,err := xorm.NewEngine("mysql",DATASOURCE)

	//是否输出查询SQL语句到终端或是日志文件里面
	//dbenginer.ShowExecTime(true)
	//dbenginer.ShowSQL(true)

	//将查询日志信息不输出到终端，而是记录到日子文件里面
	f,err := os.Create("xorm_sql.log")
	if err != nil{
		panic(err.Error())
	}
	dbenginer.SetLogger(xorm.NewSimpleLogger(f))

	if err != nil{
		panic(err.Error())
	}

	if dbenginer.Ping() != nil {
		panic(err.Error())
	}
	defer dbenginer.Close()
	tables,err := dbenginer.DBMetas()
	if err != nil{
		panic(err.Error())
	}
	for _, table := range tables {
		fmt.Printf("tabName:%s",table.Name)
	}
}
