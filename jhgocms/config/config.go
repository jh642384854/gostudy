package config

import (
	"encoding/json"
	"os"
)

//服务端配置
type AppConfig struct {
	AppName    string   `json:"app_name"`
	Port       string   `json:"port"`
	StaticPath string   `json:"static_path"`
	Mode       string   `json:"mode"`
	DataBase   DataBase `json:"data_base"`
	Redis      Redis    `json:"redis"`
}

//数据库配置信息
type DataBase struct {
	Drive        string `json:"drive"`
	Port         string `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Database     string `json:"database"`
	Params       string `json:"params"`
	ShowSql      bool   `json:"show_sql"`
	MaxOpenConns int    `json:"max_open_conns"`
	Prefix       string `json:"prefix"`
}

//Redis配置信息
type Redis struct {
	NetWork  string `json:"net_work"`
	Host     string `json:"host"`
	Databse  int    `json:"databse"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
}

func InitConfig() *AppConfig {
	var config *AppConfig

	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err.Error())
	}

	return config
}
