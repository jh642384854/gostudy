package main

import (
	"fmt"
	"net/http"
)
/**
	设置Cookie并将cookie发送到客户端
 */
func setCookie(w http.ResponseWriter,r *http.Request)  {
	c1 := http.Cookie{
		Name:"first_cookie",
		Value:"go web programming",
		HttpOnly:true,
	}
	c2 := http.Cookie{
		Name:"second_cookie",
		Value:"value2",
		HttpOnly:true,
	}
	w.Header().Set("Set-Cookie",c1.String())//注意这里使用的是Set()方法
	w.Header().Add("Set-Cookie",c2.String())//这里使用的是Add()方法。
	//或是用下面的这种方式
	//http.SetCookie(w,&c1)
	//http.SetCookie(w,&c2)
}

func getCookie(w http.ResponseWriter,r *http.Request)  {
	//直接通过Request接口的Header()方法，并获取Cookie键的值即可
	cookies := r.Header["Cookie"]
	fmt.Fprintln(w,cookies)
}

//通过Request接口的Cookie()方法来获取cookie信息
func getCookie2(w http.ResponseWriter,r *http.Request)  {
	c1,err := r.Cookie("first_cookie")
	if err != nil{
		fmt.Fprintln(w,"Cannot get the first cookie")
	}
	allCookies := r.Cookies()
	fmt.Fprintln(w,c1)
	fmt.Fprintln(w,allCookies)

}

func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/setcookie",setCookie)
	http.HandleFunc("/getcookie",getCookie)
	http.HandleFunc("/getcookie2",getCookie2)
	server.ListenAndServe()
}
