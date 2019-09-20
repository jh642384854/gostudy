package main

import (
	"fmt"
	"net/http"
)

/**
	处理器函数使用
 */
func HelloServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"hello 2")
}

func WorldServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"world 2")
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.Handle("/",http.HandlerFunc(HelloServeHTTP))
	http.HandleFunc("/hello",HelloServeHTTP)
	http.HandleFunc("/world",WorldServeHTTP)
	server.ListenAndServe()
}