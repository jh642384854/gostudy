package main

import (
	"html/template"
	"net/http"
)
//设置动作示例
func set1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template4.html")
	t.Execute(w,"hello")
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/set1",set1)
	server.ListenAndServe()
}