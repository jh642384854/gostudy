package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

/**
	ResponseWrite接口使用

 */

// http.ResponseWriter的Writer()方法
func write(w http.ResponseWriter,r *http.Request)  {
	str := "<html><head><title>go web programming</title></head><body><h1>Hello World</h1></body></html>"
	w.Write([]byte(str))
}
// http.ResponseWriter的WriteHeader()方法
func writeHeader(w http.ResponseWriter,r *http.Request)  {
	w.WriteHeader(501)
	fmt.Fprintln(w,"No such service,try next door")
}
// http.ResponseWriter的Header()方法
func header(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Location","http://www.baidu.com")
	w.WriteHeader(302)
}

//直接返回JSON格式数据
func jsonDemo(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	usermap := make(map[string]interface{})
	usermap["username"] = "zansan"
	usermap["age"] = 24
	json,_ := json.Marshal(usermap)
	w.Write(json)
}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/write",write)
	http.HandleFunc("/writeHeader",writeHeader)
	http.HandleFunc("/header",header)
	http.HandleFunc("/json",jsonDemo)
	server.ListenAndServe()
}