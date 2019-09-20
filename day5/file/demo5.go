package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
	文件的内容拷贝(简单的说就是把A文件内容拷贝到B文件中)
	下面写了两种方式：
	方式一：是用bufio.NewWriter()+bufio.WriteString()+bufio.WriteString()+bufio.Flush()的方式来拷贝内容
	方式二：是用ioutil.WriteFile()直接写入，这个更节省代码量。也适合小文件内容的拷贝
 */
func main() {
	srcFile := "D:/go2.txt"
	destFile := "D:/go3.txt"
	destFile2 := "D:/go4.txt"
	//首先是读取文件内容
	fileContent,error := ioutil.ReadFile(srcFile)
	if error != nil {
		fmt.Println("文件读取失败,err=%v",error)
	}

	//文件内容拷贝方式1
	fileHandler,error2 := os.OpenFile(destFile,os.O_WRONLY | os.O_CREATE,0666)
	defer fileHandler.Close()
	if error2 != nil {
		fmt.Println("文件读取失败,err=%v",error)
	}

	writer := bufio.NewWriter(fileHandler)
	writer.WriteString(string(fileContent))
	writer.Flush()

	//文件内容拷贝方式二：更适合小文件内容的拷贝
	err := ioutil.WriteFile(destFile2,fileContent,0666)
	if err != nil{
		fmt.Println("文件内容拷贝失败,err=%v",err)
	}
	fmt.Println("文件内容拷贝完成")
}
