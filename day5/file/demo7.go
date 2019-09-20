package main

import (
	"fmt"
	"io"
	"os"
)

/**
	拷贝文件
	拷贝文件的三种方式：https://www.jianshu.com/p/6cc1938260ba

	扩展阅读：
	我煞笔的被 bufio.Reader 小坑：https://studygolang.com/articles/6793
 */
func CopyFile3(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

func CopyFile(destFile,srcFile string) (written int64, err error) {
	//1.读取目标文件内容
	srcFileHandler,err := os.Open(srcFile)
	if err != nil{
		fmt.Println("文件读取失败,err=%v",err)
	}
	defer srcFileHandler.Close()
	//2.构建bufio.Reader对象
	//reader := bufio.NewReader(srcFileHandler)


	//3.打开或创建目标文件
	descFileHandler,err := os.OpenFile(destFile,os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil{
		fmt.Println("文件读取失败,err=%v",err)
	}
	//4.构建bufio.Writer()对象
	//write := bufio.NewWriter(descFileHandler)
	defer descFileHandler.Close()

	return io.Copy(descFileHandler,srcFileHandler) //不知道是从什么版本开始，io.Copy()
	//return io.Copy(write,reader)
}
func main() {
	srcFile := "D:/go.txt"
	destFile := "E:/go8.txt"
	_,err := CopyFile(destFile,srcFile)
	if err != nil{
		fmt.Printf("error %v",err)
	}else{
		fmt.Println("文件拷贝成功")
	}

}