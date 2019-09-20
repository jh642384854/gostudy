package structg

import "fmt"

type User struct {
	Id string
	Nickname string
	Username string
	password string
}

func GetUser() User {
	var user User
	user.Id = "100000"
	user.Nickname = "彡村扛把子"
	user.password = "fdsfdf"
	user.Username = "test003"
	return user
}

func sTest1()  {
	fmt.Println("ddd")
}