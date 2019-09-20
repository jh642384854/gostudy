package main

import (
	"fmt"
	"io/ioutil"
)

/**
	文件读取示例2
	这个例子适用于读取小文件，一次性把文件内容全部读取
	在这个例子中，并不需要打开文件句柄和关闭文件句柄，因为这个都在ioutil.ReadFile()方法中处理好了
	主要用到了ioutil.ReadFile('文件路径')方法来读取文件
	需要注意的是：需要将读取的二进制内容通过string()方法转换一下。
 */
func main() {
	filePath := "D:/go.txt"
	fileContent,error := ioutil.ReadFile(filePath)
	if error != nil{
		fmt.Println("文件读取失败,err=%v",error)
	}
	fmt.Println(fileContent)//这个读取的是二进制的内容，所以需要使用下面的方式来将二进制转化为实际内容。
	fmt.Println(string(fileContent))
}