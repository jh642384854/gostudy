package main

import (
	"fmt"
	"os/exec"
)


/**
	使用exec.Command()来执行命令并捕获命令执行的结果
 */
func main() {
	var (
		cmd2 *exec.Cmd
		err2 error
		output []byte
	)

	//cmd2 = exec.Command("C:\\Windows\\System32\\cmd.exe","go")//C:\\cygwin64\\bin\\bash.exe
	cmd2 = exec.Command("C:\\cygwin64\\bin\\bash.exe","-c","sleep 2;ls -l")
	//CombinedOutput()函数：执行命令并返回标准输出和错误输出合并的二进制切片数据。
	if output,err2 = cmd2.CombinedOutput(); err2 != nil{
		fmt.Println(err2)  //可能会出现“go exit status 127”这样的错误，这个表示就是命令找不到，这个就需要好好看下cygwin的环境配置了。注意，如果重新配置好了环境，最好把GoLand这个编辑器重启一下，让环境变量命令生效。
		return
	}
	fmt.Println(string(output))
}