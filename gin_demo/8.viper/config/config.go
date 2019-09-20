package config

type Config struct {
	Port string
	Domain string
	Mysql
	Redis
	Mq
}
//mysql的相关配置
type Mysql struct {
	Type string
	Host string
	Port string
	Name string
	User string
	Pwd string
	Prefix string
	Suffix string
	Charset string
}
//redis的相关配置
type Redis struct {
	Host string
	Port int
	Pconnect bool
	Db int
	Pwd string
	Timeout int
}
//MQ的相关配置
type Mq struct {
	Host string
	Port int
	User string
	Pwd string
	Vh string
	Topic string
}
//Mongodb的相关配置
type Mongodb struct {

}