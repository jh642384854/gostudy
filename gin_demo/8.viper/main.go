package main

import (
	"fmt"
	config2 "dev/gin_demo/8.viper/config"
	"github.com/Unknwon/goconfig"
	"github.com/spf13/viper"
	"os"
)

/**
	本例子是介绍如何使用https://github.com/spf13/viper来读取配置信息。
 */
func main() {
	// 1. 从xxx.json格式文件中获取配置
	// readFromJson()
	// 2.从xxx.yaml格式文件中获取配置
	// readFromYaml()
	// 3.从环境变量中获取数据
	// readFromEnv()
	// 4.从ini文件里面读取
	ReadFromIni()
}
/**
	从xxx.json文件中读取配置
 */
func readFromJson()  {
	var config config2.Config
	viper.SetConfigName("config") //这里不需要写文件的后缀名
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal error config file:%s\n",err))
	}
	//获取单个配置信息
	fmt.Println(viper.Get("domain"))
	//获取子级元素配置信息
	fmt.Println(viper.Get("mysql.pwd"))
	//下面调用viper.Unmarshal()方法，将读取的配置文件内容映射到定义的配置结构体上面
	if err := viper.Unmarshal(&config); err != nil{
		panic(fmt.Errorf("Fatal Unmarshal config:%s\n",err))
	}
	fmt.Println(config)
}
/**
	从xxxx.yaml配置文件中获取配置
 */
func readFromYaml()  {
	var config config2.Config
	viper.SetConfigName("config2") //这里不需要写文件的后缀名
	viper.AddConfigPath(".")
	//viper.SetConfigType("yaml") //设置配置文件类型,这里指明是yaml格式文件。这个好像并没有效果
	err := viper.ReadInConfig()
	if err != nil{
		panic(fmt.Errorf("Fatal error config file:%s\n",err))
	}
	//获取单个配置信息
	fmt.Println(viper.Get("domain"))
	//获取子级元素配置信息
	fmt.Println(viper.Get("mysql.pwd"))
	//下面调用viper.Unmarshal()方法，将读取的配置文件内容映射到定义的配置结构体上面
	if err := viper.Unmarshal(&config); err != nil{
		panic(fmt.Errorf("Fatal Unmarshal config:%s\n",err))
	}
	fmt.Println(config)
}
/**
	从环境变量中获取数据
 */
func readFromEnv()  {
	//下面是一个从环境变量里面读取数据的小例子，和我们这里要介绍的通过viper组件读取环境变量的值没有直接关联
	//fmt.Println(os.Getenv("Path"))
	if err := os.Setenv("JHMYSQLSOURCE","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=true&loc=Local"); err != nil{
		fmt.Println(err.Error())
	}
	viper.AutomaticEnv()
	//viper.SetEnvPrefix("jh")
	viper.BindEnv("mysqlsource")
	fmt.Println(viper.Get("JHMYSQLSOURCE"))
}
/**
	从xxx.ini配置文件获取配置信息
 */
func ReadFromIni()  {
	var cfg *goconfig.ConfigFile
	config,err := goconfig.LoadConfigFile("config.ini")
	if err != nil{
		fmt.Println("get config file error")
		os.Exit(-1)
	}
	cfg = config
	//获取单个值
	mysqlPwd,_ :=cfg.GetValue("sys","domain")
	fmt.Println(mysqlPwd)
	//获取全部配置
	redis,_ := cfg.GetSection("redis")
	fmt.Println(redis)
	//加载完全局配置后，该配置长驻内存，需要动态加载的话，使用cfg.Reload()方法
	if err := cfg.Reload(); err != nil{
		fmt.Printf("reload config file error: %s", err)
	}
}
