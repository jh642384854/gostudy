package main

import (
	"dev/projecttest/utils"
	"fmt"
)

func main() {
	fmt.Println("使用面向对象方式来操作")
	utils.NewAccountMange().MainMenu()
}