package main

//打开文件
/*
	注意点：
	1.使用os包的Open()方法来打开文件
	2.打开文件后，记得要关闭文件句柄。
 */

import (
	"fmt"
	"os"
)

func main() {

	filePath := "D:/go.txt"

	fileHandle,error := os.Open(filePath)
	if error != nil{
		fmt.Println("文件读取失败,err=%v",error)
	}
	//关闭文件句柄
	defer fileHandle.Close()

	//查看文件句柄
	fmt.Println(fileHandle)
}