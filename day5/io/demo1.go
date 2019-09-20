package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

/**
	io.Reader接口使用。
 */
//os.File  使用ioutil.ReadAll()方法，一次性全部读取所有数据
func demo1()  {
	file,err := os.Open("iostudy.txt")
	if err != nil{
		fmt.Println("fail open file:",err.Error())
		return
	}
	defer file.Close()
	var fileContent []byte
	fileContent,err = ioutil.ReadAll(file)  //这个读取的是二进制格式的内容
	if err != nil{
		fmt.Println("readfile fail")
	}
	fmt.Println(string(fileContent))  //使用string(fileContent)，将二进制转换为字符串格式
}
//bufio.NewReader 使用带缓存的方式来读取文件内容
func demo2()  {
	file,err := os.Open("iostudy.txt")
	if err != nil{
		fmt.Println("fail open file:",err.Error())
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	//循环读取
	for  {
		string,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Println(string)
	}
}

//bufio.Scanner 入门示例
func demo7()  {
	file,err := os.Open("iostudy.txt")
	if err != nil{
		fmt.Println("fail open file:",err.Error())
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func demo8()  {
	input := "one two three"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}



//strings.NewReader()的使用
func demo3()  {
	reader := strings.NewReader("http://studygolang.com. \nIt is the home of gophers")
	fmt.Println(reader.Size())
	//var content []byte   使用这种方式是无法把数据写入到这个变量里面的。
	//这里必须要先分配空间，因为strings.Read()方法的参数是一个切片。参数是切片，就需要为这个切片设置长度。
	// 但是如果返回值是切片，那就没有必要了，因为返回值有多大，也不确定，所以也就没有办法设置切片长度了。
	content := make([]byte,reader.Size())
	if _,err := reader.Read(content); err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

//bytes.Buffer
/**
	b1 := new(bytes.Buffer) //直接使用 new 初始化，可以直接使用
	// 其它两种定义方式
	func NewBuffer(buf []byte) *Buffer
	func NewBufferString(s string) *Buffer
 */
func demo4()  {
	buffer := bytes.NewBufferString("http://studygolang.com. \nIt is the home of gophers")
	fmt.Println(buffer.Len())
	line,err := buffer.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(line)
}

//bytes.Reader
func demo5()  {
	reader := bytes.NewReader([]byte("http://studygolang.com. \nIt is the home of gophers"))
	fmt.Println(reader.Size())
	content := make([]byte,reader.Size())
	if _,err := reader.Read(content); err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

//net.conn
func demo6()  {
	conn,err := net.Dial("tcp","127.0.0.1:8088")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	serverData := make([]byte,512)  //创建一个服务端发送给客户端数据的变量，并分配空间大小。这里固定是512字节
	if _,err := conn.Read(serverData); err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(serverData))
}


func main() {
	demo8()
}
