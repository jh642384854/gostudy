package main

import (
	"fmt"
	"errors"
	"os"
)

/**
	判断文件是否存在的方法
	主要用到了os.Stat()和os.IsNotExist()这两个方法
 */

func PathExits(path string) (bool,error){
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err){
		return  false,errors.New("file not eixts")
	}
	return false,err;
}

func main() {
	path := "D:/fdf.txt"
	hasExists,error := PathExits(path);
	if error != nil{
		fmt.Println(error)
	}
	fmt.Println(hasExists)
}
