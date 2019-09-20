package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"

	"jhgocms/config"
)

func NewMysqlEngine() *xorm.Engine {
	initConfig := config.InitConfig()
	if initConfig == nil{
		return nil
	}

	database := initConfig.DataBase

	//定义的格式如下：[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dataSourceName := database.User+":"+database.Password+"@"+"tcp("+database.Host+":"+database.Port+")/"+database.Database+"?"+database.Params
	dbEngine,err := xorm.NewEngine(database.Drive,dataSourceName)
	//dbEngine,err := xorm.NewEngine("mysql",dataSourceName)

	//定义表的前缀
	tableMapper := core.NewPrefixMapper(core.SnakeMapper{},database.Prefix)
	dbEngine.SetTableMapper(tableMapper)

	if err != nil{
		panic(err.Error())
	}

	if dbEngine.Ping() != nil {
		panic(err.Error())
	}

	//数据库相关设置
	dbEngine.ShowSQL(database.ShowSql)
	dbEngine.SetMaxOpenConns(database.MaxOpenConns)

	return  dbEngine
}