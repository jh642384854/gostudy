package main

import (
	"flag"
	"fmt"
)

/**
	使用方式：
	使用默认参数：
	PS E:\GoProjects\src\dev\day5\flag> go run .\demo2.go
	param1 = value1,param2 = 1
	使用自定义方式：注意，在参数名称前面需要输入一个短横线
	PS E:\GoProjects\src\dev\day5\flag> go run .\demo2.go -param1 VALUE1 -param2 22
	param1 = VALUE1,param2 = 22
 */

func main() {
	param1 := flag.String("param1","value1","param1 use")
	param2 := flag.Int("param2",1,"param2 use")
	//必须要执行下面的语句，对参数变量进行解析
	flag.Parse()
	//对于变量需要用*(星号，即取值符)来取值。默认返回的是一个地址
	fmt.Printf("param1 = %v,param2 = %d",*param1,*param2)
}