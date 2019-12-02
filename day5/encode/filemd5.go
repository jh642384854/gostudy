package main

import (
	"fmt"
	"crypto/md5"
	"io"
	"os"
)

func main() {

	filePath := "./uploads/20191021/2DXo9gJHbd.png"
	file,err := os.Open(filePath)
	if err == nil{
		md5h := md5.New()
		io.Copy(md5h, file)
		fmt.Printf("%x", md5h.Sum([]byte(""))) //md5
		fmt.Println()
		fmt.Println(fmt.Sprintf("%x",md5h.Sum([]byte(""))))
	}else{
		fmt.Println(err.Error())
	}

}