package main


import (
	"flag"
	"fmt"
)

/**
	使用flag包来解析命令行参数

	上面的一个例子虽然可以获取可执行文件传递的参数，但是每个参数传递的时候并没有太大意义，根本不知道每个参数代表什么意思
	这里就用另外一种方式来实现。使用flag包来解析命令行参数。
	就是类似mysql -u root -h 127.0.0.1 -p xxxx这样的，我们可以来解析-u -h -p等等这样的参数值。
	下面示例如下：
	PS E:\GoProjects\src\dev\day5\cmd> .\demo2.exe -u jiang -pwd 154575
	user=jiang,pwd=154575,host=127.0.0.1,port=123456
	注意：
	在执行可执行命令的时候，所携带的参数必须以"-u"这样的格式，而不能只是用"u"(没有带段横线)，这样是无法解析参数的。
 */

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","密码，默认为空")
	flag.StringVar(&host,"h","127.0.0.1","主机，默认为127.0.0.1")
	flag.IntVar(&port,"p",123456,"密码，默认为123456")

	//这里有一个非常重要的操作，就是调用flag.Parse()进行转换。
	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v",user,pwd,host,port)
}
