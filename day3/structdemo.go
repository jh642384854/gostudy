package main

import (
	"dev/day3/structg"
	"fmt"
)

func main() {
	structg.TestSendMail()
	structg.TestGetUser()

	username := "zhangsan"
	var name *string = &username
	fmt.Println(username,name)
}
