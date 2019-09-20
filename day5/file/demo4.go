package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
	文件的写操作
	os.Open()和os.OpenFile()的区别和共同点：
	共同点：都可以打开文件
	区别：os.Open()只能打开文件，而os.OpenFile()除了可以打开文件，还可以对打开的文件进行操作(重新写入等等)

	用到了os.OpenFile()来打开操作的文件
	用到了bufio.NewWriter()来创建一个bufio.Writer对象，用来写入内容
	用到了bufio.WriteString()方法来执行写入操作(这里的写入是写入到内存，而并不是写入到磁盘)
	用到了bufio.Flush()方法，这个方法就是将文件真正的写入到磁盘上面。
 */
func main() {
	filePath := "D:/go2.txt"
	fileHandler,error := os.OpenFile(filePath,os.O_WRONLY | os.O_CREATE,0666);//这里关键点就是os.OpenFile()的第二个参数，用来指定对打开的文件进行何种模式的操作。具体可以查看os.file包中定义的常量
	if error != nil{
		fmt.Println("创建文件失败,err=%v",error)
	}
	defer fileHandler.Close()
	str := "go study \r\n"  //这里为什么是\r\n?  因为在不同的编辑器下面，对换行的解析不太一样。比如记事本是用\r来表示换行，而notepad++而是用的\n来表示换行
	writer := bufio.NewWriter(fileHandler)
	for i:=0;i<5 ;i++  {
		writer.WriteString(str)
	}
	writer.Flush() //这一步非常重要。是真正把内容写入到磁盘的操作
	fmt.Println("文件写入成功")
}
