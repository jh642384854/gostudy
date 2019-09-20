package main

import (
	"fmt"
	"reflect"
	"net/http"
	"runtime"
	"time"
)

/**
	串联多个处理器和处理器函数
 */
//形式一：
func hello(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"hello 77777")
}

func log(h http.HandlerFunc) http.HandlerFunc  {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - "+name)
		h(writer,request)
	}
}
//形式二:
func world(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"world 77777")
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		fmt.Println(timeElapsed)
	})
}
func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/hello",log(hello)) //一个是用的处理器函数
	http.Handle("/world",timeMiddleware(http.HandlerFunc(world)))  //一个是用的处理器
	server.ListenAndServe()
}
