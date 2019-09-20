package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
	当有多个任务在执行同一个操作，我们只需要其中一个任务执行完毕后，就可以直接通过其中的任意一个返回结果来返回

 */

func main() {
	//输出当前系统中的协程数
	fmt.Println("Before", runtime.NumGoroutine())

	fmt.Println(FirstResponse())

	time.Sleep(time.Second)

	fmt.Println("After", runtime.NumGoroutine())
}

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	/**
		在这个示例中，如果使用的是不带buffer的channel，通过runtime.NumGoroutine()函数，我们可以在执行该方法的前后会发现会有很多没有被处理的协程。
		如果在系统中又很多这样的限制协程未被处理，对系统是很不好的，有可能还会导致内存被占满的情况。
		这是为什么呢？
		我们知道不带buffer的channel，当往里面写数据的时候，如果没有读取出来，就会被出于阻塞状态，而带buffer的channel不会有这个问题。

		解决这个问题，也很简单，我们创建一个带buffer的channel就好了。
	 */
	//ch := make(chan string)   //定义为这样的channel在执行该代码的时候，会有很多没有被处理的channel存在，会造成内存泄漏的问题
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(num int) {
			res := runTask(num)
			ch <- res
		}(i)
	}
	return <-ch
}
