package main

import (
	"fmt"
	"github.com/spf13/afero"
	"os"
)

var (
	appFs = afero.NewMemMapFs()
	//appFs2 = afero.NewOsFs()
)

func main() {
	//demo1()
	//demo2()
	demo3()
}

func demo3()  {
	if err := os.MkdirAll("jhgostatic/20191018/16",os.ModePerm); err != nil{
		fmt.Println(err.Error())
	}
}

func demo2()  {
	file,err := appFs.Create("jhtest.txt")
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(file)
	if err := appFs.Mkdir("statics",os.ModePerm); err != nil{
		fmt.Println(err.Error())
	}
}

func demo1()  {
	fmt.Println(afero.DirExists(appFs,"./upload"))
}
