package main

import (
	"fmt"
	"sync"
)

/**
	使用锁来解决goroutine并行执行出现的资源竞争问题
	如果不用lock机制，在使用goroutine处理并行操作的时候，可能会出现“fatal error: concurrent map writes”这个问题。

	我们在检查一个业务是否存在资源竞争的问题，我们在编译代码的时候，可以添加-race选项。如下所示：
	PS E:\GoProjects\src\dev\day5\goroutine> go build --race .\demo2.go
	然后我们执行编译后的文件：
	PS E:\GoProjects\src\dev\day5\goroutine> .\demo2.exe
	==================
	WARNING: DATA RACE
	Read at 0x00c000056240 by main goroutine:
	  runtime.mapiterinit()
		  D:/GoEnv/src/runtime/map.go:734 +0x0
	  main.main()
		  E:/GoProjects/src/dev/day5/goroutine/demo2.go:34 +0xd5

	Previous write at 0x00c000056240 by goroutine 40:
	  runtime.mapassign_fast64()
		  D:/GoEnv/src/runtime/map_fast64.go:92 +0x0
	  main.cal()
		  E:/GoProjects/src/dev/day5/goroutine/demo2.go:24 +0xa1

	Goroutine 40 (finished) created at:
	  main.main()
		  E:/GoProjects/src/dev/day5/goroutine/demo2.go:31 +0x6c
	==================
	nums[12] = 479001600
	nums[25] = 7034535277573963776
	........
	Found 2 data race(s)
	就会发现这样的问题。如果goroutine出现了资源竞争的问题，就可以使用lock的机制来处理。但是更加推荐使用channel方式来处理。
	因为使用lock，你不仅需要在写入数据的地方考虑加锁和解锁的操作，在读取数据的时候，也一样需要，这样灵活性还是不够。代码的依赖也较强
 */
var (
	nums = make(map[int]int)
	lock sync.Mutex
)
func cal(n int)  {
	res := 1
	for i := 1;i<=n;i++  {
		res *=i
	}
	//

	lock.Lock()
	nums[n] = res
	lock.Unlock()
}

func main() {

	for i:= 0;i<=100 ;i++  {
		go cal(i)
	}
	lock.Lock()
	for index,v := range nums{
		fmt.Printf("nums[%v] = %v \n",index,v)
	}
	lock.Unlock()
}