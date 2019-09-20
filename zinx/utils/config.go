package utils

import (
	"encoding/json"
	"io/ioutil"
)

type GloabConfig struct {
	ServerName        string
	Host              string
	Port              int
	MaxBuffer         uint32
	Version           string
	MaxConnectionSize int    //最大链接数
	MaxPackageSize    uint32 //最大的发包数据量
	MaxWorkPoolSize   uint32 //最大的工作池数量
	MaxTaskQueueLen   uint32 //每个工作池中队列里面保存的最大数据量
}

var GloabConfigObj *GloabConfig

func (this *GloabConfig) ReloadConfig() {
	config, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}
	//将解析的json转换为对应的结构体
	err = json.Unmarshal(config, &GloabConfigObj)
	if err != nil {
		panic(err)
	}
}

func init() {
	GloabConfigObj = &GloabConfig{
		ServerName:        "zinx Server",
		Host:              "0.0.0.0",
		Port:              8989,
		MaxBuffer:         1024,
		Version:           "v0.4",
		MaxConnectionSize: 1,
		MaxPackageSize:    4096,
		MaxWorkPoolSize:   10,
		MaxTaskQueueLen:   1024,
	}
	//GloabConfigObj.ReloadConfig()
}
