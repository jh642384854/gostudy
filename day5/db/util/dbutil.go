package util

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"sync"
)

var singletonDB *sql.DB
var once sync.Once

func GetDbInstance() *sql.DB {
	once.Do(func() {
		db,err := sql.Open("mysql","root:jianghua@tcp(localhost:3306)/godb?charset=utf8")
		if err != nil{
			log.Fatal(err)
		}
		/*if err = db.Ping(); err != nil{
			log.Fatal(err)
		}*/
		singletonDB = db
	})
	return singletonDB
}

