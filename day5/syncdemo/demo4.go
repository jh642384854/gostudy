package main

import (
	"fmt"
)

func main() {
	done := make(chan int)

	go func() {
		fmt.Println("hello")
		done <- 1
	}()
	/**
		当<-done执行时，必然要求done <- 1也已经执行。
		根据同一个Gorouine依然满足顺序一致性规则，我们可以判断当done <- 1执行时，println("你好, 世界")语句必然已经执行完成了。因此，现在的程序确保可以正常打印结果。
	 */
	<- done
}
