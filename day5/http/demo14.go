package main

import (
	"html/template"
	"net/http"
)
//包含动作示例
func include1(w http.ResponseWriter,r *http.Request)  {
	//在处理函数中，我们必须将所有模板文件都进行语法分析。需要传入多个文件，其中，第一个模板文件是主模板，也就是说，主模板必须放第一个。
	t := template.Must(template.ParseFiles("E:/GoProjects/src/dev/day5/http/template5.html","E:/GoProjects/src/dev/day5/http/template_include.html"))
	t.Execute(w,"hello")
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/include1",include1)
	server.ListenAndServe()
}