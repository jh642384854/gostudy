package main
//文件的读取示例
/**
	用到了io.Open('文件路径')打开文件句柄
	用到了bufio.NewReader('打开的文件句柄')来创建文件读取句柄
	用到了bufio.ReadString('分割符号')来读取文件内容
	用到了io.EOF来判断是否读取文件完毕
 */
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//读取文件
func main() {

	filePath := "D:/go.txt"
	//1.打开一个文件句柄
	fileHandle,error := os.Open(filePath)
	if error != nil{
		fmt.Println("文件读取失败,err=%v",error)
	}
	//2.关闭文件句柄
	defer fileHandle.Close()

	//3.依据打开的文件句柄来创建一个bufio.Reader对象，根据这个对象来读取文件内容
	readHandler := bufio.NewReader(fileHandle)
	//循环读取文件内容
	for {
		//4.逐行来读取文件内容
		str,error := readHandler.ReadString('\n') //
		//5.判断文件是否读取到文件末尾，如果读取结束，就退出循环体
		if error == io.EOF{
			break
		}
		//6.打印读取的文件内容
		fmt.Println(str)
	}
	fmt.Println("文件已经读取完成")

}
