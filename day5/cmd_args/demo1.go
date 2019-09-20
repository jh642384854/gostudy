package main
/**
	通过os.Args属性来获取执行编译后的可执行文件传递的参数
	示例如下：
	PS E:\GoProjects\src\dev\day5\cmd> .\demo1.exe jkj kjjfks 234
	获取的参数个数： 4
	arg[0] = E:\GoProjects\src\dev\day5\cmd\demo1.exe
	arg[1] = jkj
	arg[2] = kjjfks
	arg[3] = 234
 */
import (
	"fmt"
	"os"
)

func main() {

	osArgs := os.Args
	//os.Args得到的是一个切片
	fmt.Println("获取的参数个数：",len(osArgs))

	//遍历os.Args切片数据
	for index,v := range osArgs {
		fmt.Printf("arg[%v] = %v \n",index,v)
	}
}
