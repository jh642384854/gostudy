package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
	数据库连接demo
 */
func main() {
	// 创建数据库对象需要引入标准库database/sql，同时还需要引入驱动go-sql-driver/mysql。
	// 使用_表示引入驱动的变量，这样做的目的是为了在你的代码中不至于和标注库的函数变量namespace冲突。
	db,err := sql.Open("mysql","root:数据连接密码@tcp(localhost:3306)/godb?charset=utf8")
	if err != nil{
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()
}
