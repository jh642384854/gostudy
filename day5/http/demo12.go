package main

import (
	"html/template"
	"net/http"
)

//迭代动作入门
func range1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template3.html")
	skills := []string{"php","java","go"}
	t.Execute(w,skills)
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/range1",range1)
	server.ListenAndServe()
}