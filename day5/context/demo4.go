package main

import (
	"context"
	"time"
	"fmt"
)
var key string = "name"
func main() {
	ctx,cancel := context.WithCancel(context.Background())

	valueCtx := context.WithValue(ctx,key,"【监控1】")
	go watch2(valueCtx)

	valueCtx = context.WithValue(ctx,key,"【监控2】")
	go watch2(valueCtx)

	valueCtx = context.WithValue(ctx,key,"【监控3】")
	go watch2(valueCtx)

	time.Sleep(5*time.Second)
	fmt.Println("可以了，通知监控停止")
	defer cancel()
	time.Sleep(5*time.Second)
}

func watch2(ctx context.Context)  {
	for  {
		select {
		case <- ctx.Done():
			fmt.Println(ctx.Value(key),"监控退出，停止了...",ctx.Err())
			return
		default:
			fmt.Println("go routine持续监控中...")
			time.Sleep(2*time.Second)
		}
	}
}