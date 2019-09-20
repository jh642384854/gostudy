package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan int)

	go func() {
		c1 <- 1
	}()


	select {
	case v := <-c1:
		fmt.Println(v)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}

}
