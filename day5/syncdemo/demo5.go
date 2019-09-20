package main

import (
	"fmt"
	"sync"
)

/**
	sync.Map的使用
 */

func demo1()  {
	//声明一个sync.Map的变量
	var sm sync.Map
	//使用Store()方法，添加元素
	sm.Store(1,"a")
	//Load()方法，获取value
	if val,ok := sm.Load(1);ok{
		fmt.Println(val)
	}
	//LoadOrStore()方法，用来获取或保存数据
	//参数是一对key:val。如果该key存在且没有被标记删除则返回原先的value(即不进行更新)和true;如果不存在，则将会报错，返回该value和false
	if val2,ok := sm.LoadOrStore(1,"c"); ok{
		fmt.Println(val2)
	}
	if val3,ok := sm.LoadOrStore(2,"c");ok{
		fmt.Println(val3)
	}else{
		fmt.Println("new val:",val3)
	}
	//遍历sync.Map里面存储的数据
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key,":",value)
		return true
	})
}

func main() {
	demo1()
}