package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	const dataFile = "../data/data.json"
	_, filename, _, _ := runtime.Caller(1)
	datapath := path.Join(path.Dir(filename), dataFile)
	fmt.Println(datapath)
	execpath, _ := os.Executable() // 获得程序路径
	fmt.Println(execpath)
	//configfile := filepath.Join(filepath.Dir(execpath), "./config.yml")
}
