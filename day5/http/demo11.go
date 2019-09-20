package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"html/template"
)

//条件动作入门

func condition1(w http.ResponseWriter,r *http.Request)  {
	t,_ := template.ParseFiles("E:/GoProjects/src/dev/day5/http/template2.html")
	rand.Seed(time.Now().Unix())
	num := rand.Intn(10)
	fmt.Println(num)
	t.Execute(w,num>5)
}


func main() {
	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/condition1",condition1)
	server.ListenAndServe()
}