package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)
/**
	自定义函数
	将自定义函数应用到模板中
 */
func formatDate(t time.Time) string {
	layout := "2006-01-02 15:04:05"  //年、月、日、时、分、秒，英文、数字，必须精确地限定到 golang 指定的时间原点：2006-01-02 15:04:05
	return t.Format(layout)
}

func diyfun(w http.ResponseWriter,r *http.Request)  {
	funcMap := template.FuncMap{"fdate":formatDate}
	templateFile := "E:/GoProjects/src/dev/day5/http/template6.html"
	//template.New() 模板名要用文件名，不能是带路径的名字。不然在执行程序的时候，就会提示template: “…” is an incomplete or empty template这样的错误
	t := template.New(path.Base(templateFile)).Funcs(funcMap)
	t,err := t.ParseFiles(templateFile)
	if err != nil{
		fmt.Println(err)
	}
	err = t.Execute(w,time.Now())
	fmt.Println(err)

}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/diyfun",diyfun)
	server.ListenAndServe()
}