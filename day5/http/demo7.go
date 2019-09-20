package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
	获取表单提交数据
	文件上传保存到指定目录：https://www.jianshu.com/p/5f29ef2daf55
 */

func process(wr http.ResponseWriter,r *http.Request)  {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file , err := fileHeader.Open()
	if err == nil{
		data,err := ioutil.ReadAll(file)
		if err == nil{
			fmt.Fprintln(wr,string(data))
		}
	}
}
//处理单个文件上传
func processSignal(w http.ResponseWriter,r *http.Request)  {
	file,_,err := r.FormFile("uploaded")
	if err == nil{
		data,err := ioutil.ReadAll(file)
		if err == nil{
			fmt.Fprintln(w,string(data))
		}
	}
}

func main() {

	server := http.Server{
		Addr:":8080",
	}
	http.HandleFunc("/process",process)

	server.ListenAndServe()

}