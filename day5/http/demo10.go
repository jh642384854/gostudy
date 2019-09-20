package main

import (
	"html/template"
	"net/http"
)
//模版入门示例

//针对的模版文件
func tem1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template1.html")
	t.Execute(w,"hello world")
}

//针对的是模版字符串
func tem2(w http.ResponseWriter,r *http.Request)  {
	tmp1 := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
{{ . }}
</body>
</html>`
	t := template.New(tmp1)
	t,_ = t.Parse(tmp1)
	t.Execute(w,"hello string")
}

//template.Must()方法的使用：用来检查进行模版语法分析时候是否会报错
func tem3(w http.ResponseWriter,r *http.Request)  {
	//t := template.Must(template.ParseFiles("E:/GoProjects/src/dev/day5/http/template1.html"))
}

func main() {

	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/tem1",tem1)
	http.HandleFunc("/tem2",tem2)
	server.ListenAndServe()
}