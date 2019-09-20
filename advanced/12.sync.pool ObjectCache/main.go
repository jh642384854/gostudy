package main

import (
	"fmt"
	"sync"
	"runtime"
)

/**
	对象缓存

	尝试从私有对象获取
	私有对象不存在的时候，就会尝试从当前Processor的共享池中获取
	如果当前Processor共享池也是空的，那么就会尝试去其他Processor的共享池获取
	如果所有的池都是空的，最后就用用户指定的New函数产生一个新的对象返回。

	单个私有对象：协程安全
	共享池(多个对象存放地)：协程不安全

	sync.Pool对象的放回
	如果私有对象不存在，则保存为私有对象
	如果私有对象存在，则放入当前processor子池的共享池中

	sync.Pool总结
	适用于通过复用，降低复杂对象的创建和GC代价
	协程安全，会有锁的开销
	生命周期会收到GC的影响，不太适合做连接池等，需要自己管理生命周期
 */

func main() {
	TestSyncPool()
}

func TestSyncPool()  {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)

	//手动调用GC处理机制
	runtime.GC()

	v1,_ := pool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine()  {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
		}(i)
	}
	wg.Wait()
}