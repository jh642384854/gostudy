package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd *exec.Cmd
	err error
)
/**
	使用exec.Command()来执行命令
 */
func main() {
	//生成CMD
	cmd = exec.Command("C:/cygwin64/bin/bash.exe","-c","echo 1")
	err = cmd.Run()
	fmt.Println(err)
}