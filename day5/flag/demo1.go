package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice"
	if len(os.Args) > 1{
		//这里为什么是os.Args[1:]？
		//因为os.Args[0]表示的是当前脚本执行的命令路径。而os.Args[1:]所表示的就是不要包含命令路径，只获取命令后面的参数
		who += strings.Join(os.Args[1:]," ")
	}
	fmt.Println("Good Morning",who)
}
