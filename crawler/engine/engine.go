package engine

import (
	"dev/crawler/fetcher"
	"fmt"
	"log"
)
/**
	非并发的处理方式
 */

type SimpleEngine struct { }

func (s SimpleEngine) Run(seeds ...Request)  {
	var requests []Request  //定义的一个队列，这里使用的是用数组作为队列
	for _, r := range seeds {
		requests = append(requests,r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]
		fmt.Println("当前追踪的URL是：",r.Url)
		parseResult,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests,parseResult.Requests...)
		//fmt.Println(requests)
		for _, item := range parseResult.Items {
			log.Printf("got item %v，url:%s \n",item,r)
		}
	}
}

func worker(r Request) (ParseResult,error) {
	body,err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher Error:fetcher url %s %v",r.Url,err)
		return ParseResult{},fmt.Errorf("Fetcher Error:fetcher url %s %v",r.Url,err)
	}
	return r.ParserFun(body),nil
}
