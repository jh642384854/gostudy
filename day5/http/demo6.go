package main

import (
	"fmt"
	"net/http"
)

/**
	获取请求首部信息
 */
func headers(wr http.ResponseWriter,r *http.Request)  {
	headers := r.Header
	fmt.Fprintln(wr,headers)
}

/**
	获取请求主体信息
 */
func body(wr http.ResponseWriter,r *http.Request)  {
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	fmt.Fprintln(wr,string(body))
}

/**
	parseForm
 */

func parseForm(wr http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Fprintln(wr,r.Form)                  // 存储了post、put和get参数
	fmt.Fprintln(wr,r.PostForm)              // 存储了post、put参数
	fmt.Fprintln(wr,r.FormValue("username")) //获取指定key的参数
	fmt.Fprintln(wr,r.PostFormValue("age")) //获取指定key的参数
}

func main() {

	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/headers",headers)
	http.HandleFunc("/body",body)
	http.HandleFunc("/form",parseForm)

	server.ListenAndServe()
}
