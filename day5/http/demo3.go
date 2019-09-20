package main

import (
	"fmt"
	"net/http"
)

/**
	处理器使用
 */


type HelloHandle struct {}
type WorldHandle struct {}

func (h *HelloHandle) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"hello")
}

func (h *WorldHandle) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintln(w,"world")
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.Handle("/hello",&HelloHandle{})
	http.Handle("/world",&WorldHandle{})

	server.ListenAndServe()

}