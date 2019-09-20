package main

import (
	"fmt"
	"net/http"
)


type MyHandler struct {

}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world")
}

func main() {
	httpServer := &http.Server{
		Addr:":8080",
		Handler:&MyHandler{},   //用自定义的Handle来替换默认的DefaultServeMux。
	}
	httpServer.ListenAndServe()
}