package main

import (
	"fmt"
	"sync"
	"unsafe"
)

/**
	本示例主要演示了sync.Once的使用
 */

func main() {
	var wg sync.WaitGroup
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go func() {
			obj := GetSingleObj()
			//unsafe.Pointer()用来获取对象的地址，下面输出的结果都是一样的，既然对象的内存地址一样，那就说明肯定是同一个对象了
			fmt.Println(unsafe.Pointer(obj))
			//fmt.Printf("%x\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

type SingleObj struct {

}
var singleObjInstance *SingleObj
var once sync.Once

//创建单一对象(就是单例模式)
func GetSingleObj() *SingleObj {
	once.Do(func() {
		fmt.Println("created SingleObj")
		singleObjInstance = new(SingleObj)
	})
	return singleObjInstance
}



