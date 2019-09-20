package main

import (
	"fmt"
	"reflect"
)

type User struct {
	UserName string
	Age int
}

func (this User) NewUser(name string,age int) *User {
	return &User{
		UserName:name,
		Age:age,
	}
}

func (this User) GetUserInfo()  {
	fmt.Println("get Userinfo")
}

func (this User) DeleteUser()  {
	fmt.Println("delete user")
}

func main() {

	user := User{"wangwu",25}

	v := reflect.ValueOf(user)
	t := reflect.TypeOf(user)

	fileds := v.NumField()
	methods := v.NumMethod()
	fmt.Printf("User struct 有%v个属性,有%v个方法 \n",fileds,methods);

	//遍历属性
	for i:=0; i<fileds ;i++  {
		filedName := t.Field(i)
		filedValue := v.Field(i)
		fmt.Printf("[%v] = %v \n",filedName.Name,filedValue.Interface())
	}





}
