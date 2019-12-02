package main

import (
	"fmt"
	"path"
)

func main() {
	fullFilename := "/Users/itfanr/Documents/test.txt"
	//获取带文件名的后缀
	filenameWithSuffix := path.Base(fullFilename)
	fmt.Println(filenameWithSuffix)
	//直接获取文件的后缀
	fileSuffix := path.Ext(filenameWithSuffix)
	fmt.Println(fileSuffix)
	fmt.Println(fileSuffix[1:len(fileSuffix)])
}