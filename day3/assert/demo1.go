package main

import "fmt"

type User struct {
	Name string
}

func main() {

	var user interface{}
	user2 := User{
		Name:"zhangsan",
	}
	user = user2
	fmt.Println(user)

	var user3 User       //这里先声明一个类型为User的变量user3
	//user3 = user       //把声明为空接口的变量user赋值给user3，如果单纯的这样写市不对的，会有类型转换失败的问题
	user3 = user.(User)  //换成这样的写法，这里就是用到的类型断言.

	user4 := user   //这个有别于上面var user3 User这种声明方式

	fmt.Println(user3,user4)




}
