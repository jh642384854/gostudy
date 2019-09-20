package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
func echo(wr http.ResponseWriter,r *http.Request)  {
	msg,err := ioutil.ReadAll(r.Body)
	if err != nil{
		wr.Write([]byte("echo error"))
		return
	}
	writelen,err := wr.Write(msg)
	if err != nil || writelen != len(msg){
		log.Println(err,"write len:",writelen)
	}
}


func main() {

	http.HandleFunc("/",echo)
	//第一种创建http Server的方式，这是一种快捷方式，只能传递2个参数，分别是服务器绑定的端口和handler(处理器)
	//注意下面监听地址的写法，是:9100，端口号前面还有一个冒号。
	err := http.ListenAndServe(":9100",nil)
	if err != nil{
		log.Fatal(err)
	}
	//第二种创建http Server的方式。这种方式能对http Server进行更详细的配置，具体可以参考http/Server的结构体定义
/*
	httpServer := http.Server{
		Addr:":8080",
		Handler:nil,
	}
	err := httpServer.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
*/
}


