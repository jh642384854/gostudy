package main

import (
	"fmt"

	"github.com/go-redis/redis"

)

func main() {

	client := InitRedis()

	defer client.Close()

	strCmd := client.Get("name")
	fmt.Println(strCmd.Val())
}

func InitRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:"localhost:6379",
		Password:"",
		DB:0,
	})

	pong,err := redisClient.Ping().Result()
	if err != nil{
		panic(err.Error())
	}
	fmt.Println(pong)

	return redisClient
}