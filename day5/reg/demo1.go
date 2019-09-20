package main

import (
	"fmt"
	"regexp"
)

var emailStr = `first email is abc@qq.com
second email is def@163.com
third email is gh@sina.com.cn`

func main() {

	regExpression := `[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`
	regExpression2 := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`

	reg := regexp.MustCompile(regExpression)     //创建一个正则表达式对象
	emails := reg.FindAllString(emailStr,-1)  //获取匹配的结果
	fmt.Println(emails)

	reg2 := regexp.MustCompile(regExpression2)
	emails2 := reg2.FindAllStringSubmatch(emailStr,-1)
	fmt.Println(emails2)
}
