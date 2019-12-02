package main

import (
	"fmt"
	"github.com/allegro/bigcache"
	"log"
	"time"
)

func main() {
	bigcacheConfig := bigcache.Config{
		Shards:2,// 存储的条目数量，值必须是2的幂
		LifeWindow:10*time.Minute,
	}
	cache,err := bigcache.NewBigCache(bigcacheConfig)
	if err != nil{
		log.Fatal(err)
	}
	//设置值
	cache.Set("username",[]byte("zhangsan"))
	//获取值
	username,_ := cache.Get("username")
	fmt.Println(string(username))
}
