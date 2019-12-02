package main

import (
	"fmt"
	"github.com/jinzhu/now"
)
/**
	github.com/jinzhu/now组件的使用
 */
func main() {
	fmt.Println(now.BeginningOfMonth())
	fmt.Println(now.EndOfMonth())
	fmt.Println(now.MustParse("2016"))
}