package main

import (
	"dev/crawler/engine"
	"dev/crawler/scheduler"
	"dev/crawler/storage"
	"dev/crawler/zhenai/parser"
)

/**
项目作者地址：https://github.com/liabio/crawler
 */

const (
	CityUrl = "http://www.zhenai.com/zhenghun"
)

func main() {
	/*
	engine.SimpleEngine{}.Run(engine.Request{
		Url:CityUrl,
		ParserFun:parser.ParseCityList,
	})
	*/
	/*
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:10,
	}
	e.Run(engine.Request{
		Url:CityUrl,
		ParserFun:parser.ParseCityList,
	})
	*/
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    storage.SaveItem(),
	}
	e.Run(engine.Request{
		Url:       CityUrl,
		ParserFun: parser.ParseCityList,
	})
}
