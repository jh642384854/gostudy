package main

import (
	"html/template"
	"net/http"
)
/**
	go模板上下文感知入门示例
 */
func context1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template7.html")
	content := `I asked : <i>"What's up?'"</i>`
	t.Execute(w,content)
}
/**
	go模板上下文感知的XSS处理
 */
func xss(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template8.html")
	content := `<script>alert("hello world")</script>`
	t.Execute(w,content)
}

func unxss(w http.ResponseWriter,r *http.Request)  {
	//发送一个最初由微软为IE浏览器创建的特殊HTTP响应首部X-XSS-Protection来让浏览器关闭内置的XSS防御功能
	w.Header().Set("X-XSS-Protection","0")
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template8.html")
	content := `<script>alert("hello world")</script>`
	t.Execute(w,template.HTML(content))
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/context1",context1)
	http.HandleFunc("/xss",xss)
	http.HandleFunc("/unxss",unxss)
	server.ListenAndServe()
}