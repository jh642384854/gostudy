package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)
/**
	嵌套模版的使用
 */
func layout1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template9_layout1.html")
	//Template结构体的ExecuteTemplate()方法，第二个参数接受的是一个布局名称，这个布局名称在上面的代码解析中有定义。
	t.ExecuteTemplate(w,"layout","")
}
//应用不同模板中定义的content定义动作
func layout2(w http.ResponseWriter,r *http.Request)  {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5{
		t,_ = template.ParseFiles("E:/GoProjects/src/dev/day5/http/template10_layout1.html","E:/GoProjects/src/dev/day5/http/red_hello.html")
	}else{
		t,_ = template.ParseFiles("E:/GoProjects/src/dev/day5/http/template10_layout1.html","E:/GoProjects/src/dev/day5/http/blue_hello.html")
	}
	t.ExecuteTemplate(w,"layout","")
}
//使用block动作来定义默认模板
func layout3(w http.ResponseWriter,r *http.Request)  {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5{
		t,_ = template.ParseFiles("E:/GoProjects/src/dev/day5/http/template11_layout1.html","E:/GoProjects/src/dev/day5/http/red_hello.html")
	}else{
		t,_ = template.ParseFiles("E:/GoProjects/src/dev/day5/http/template11_layout1.html")
	}
	t.ExecuteTemplate(w,"layout","")
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/layout1",layout1)
	http.HandleFunc("/layout2",layout2)
	http.HandleFunc("/layout3",layout3)
	server.ListenAndServe()
}