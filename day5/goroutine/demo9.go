package main

import "fmt"

/**
	gorouteine的异常捕获处理
 */

//下面这个方法是正常运行的
func GoTest1()  {
	fmt.Println("GoTest1")
}

//下面这个方法是会有问题的，因为map类型需要先make才能被使用。
func GoTest2()  {
	defer func() {
		if err:= recover(); err != nil {
			fmt.Println("GoTest2()方法运行时出错了")
		}
	}()
	var userMap map[string]string
	//userMap = make(map[string]string) //故意注释这行，就是为了执行上面的defer方法，用来捕获错误信息，进而不会影响整个程序出问题
	userMap["username"] = "wangwu"
}

func main() {

	go GoTest1()

	go GoTest2()

	fmt.Println("主线程运行结束")
}