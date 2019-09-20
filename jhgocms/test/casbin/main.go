package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"jhgocms/datasource"
)

func main() {
	//mysqlAdapter,err := xormadapter.NewAdapter("mysql","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8",true)
	engine := datasource.NewMysqlEngine()
	mysqlAdapter,err := xormadapter.NewAdapterByEngine(engine)
	if err != nil{
		panic(err.Error())
		return
	}

	enforcer,err := casbin.NewEnforcer("./rbac_model.conf",mysqlAdapter)

	enforcer.EnableLog(true)

	if err != nil{
		panic(err.Error())
		return
	}

	if err = enforcer.LoadPolicy(); err != nil{
		panic(err.Error())
	}



	resBool,err := enforcer.Enforce("alice", "data1", "read")
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(resBool)

	//enforcer.AddPolicy("alice", "data1", "write")

	if err = enforcer.SavePolicy(); err != nil{
		panic(err.Error())
	}

}
