package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
/*	md5Hash := md5.New()
	md5Hash.Write([]byte("12346"))
	fmt.Println(string(md5Hash.Sum(nil)))
*/
	w := md5.New()
	io.WriteString(w,"123456") //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	fmt.Println(md5str2)

	originStr := "123456"
	data := []byte(originStr)
	md5hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x",md5hash)
	fmt.Println(md5str)
}


