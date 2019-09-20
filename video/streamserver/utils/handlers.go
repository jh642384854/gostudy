package utils

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
//读取视频内容(这个就是利用浏览器的播放视频功能，并没有使用任何播放器)
func StreamHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	vid := p.ByName("video-id")
	videoUrl := VIDEO_DIR + vid

	if !strings.HasSuffix(videoUrl,"mp4") {
		videoUrl = videoUrl + ".mp4"
	}

	video,err := os.Open(videoUrl)
	if err != nil{
		log.Printf("Error when try to open file :%v",err)
		SendErrorResponse(w,http.StatusInternalServerError,"Fail open file")
		return
	}
	defer video.Close()

	w.Header().Set("Content-Type","video/mp4")
	http.ServeContent(w,r,"",time.Now(),video)
}

//上传视频
func UploadHandler(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//判断文件上传的大小
	/**
		MaxBytesReader类似于io.LimitReader但被设计用来限定接收的请求体的大小。
		不同于io.LimitReader，MaxBytesReader的返回值是一个ReadCloser，当读取超过限制时会返回non-EOF错误并且当它的关闭方法调用时会关闭潜在的读取者（函数/进程）。
		MaxBytesReader保护客户端避免偶然或者恶意发送的长数据请求导致的server资源的浪费。
	 */
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE);err != nil{
		log.Printf("Upload file to big")
		SendErrorResponse(w,http.StatusInternalServerError,"Upload file to big")
		return
	}

	//从html的表单控件来获取上传文件
	file,_,err := r.FormFile("file")
	if err != nil{
		log.Printf("Get file error:%v",err)
		SendErrorResponse(w,http.StatusInternalServerError,"Get file error")
		return
	}

	//将HTML表单元素的文件内容读取出来
	data,err := ioutil.ReadAll(file)
	if err != nil{
		log.Printf("Read file error:%v",err)
		SendErrorResponse(w,http.StatusInternalServerError,"Read file error")
		return
	}

	//将读取的文件进行保存
	fileName := p.ByName("video-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName,data,0666)
	if err != nil{
		log.Printf("Writer file error:%v",err)
		SendErrorResponse(w,http.StatusInternalServerError,"Writer file error")
		return
	}
	//返回成功信息
	w.WriteHeader(200)
	io.WriteString(w,"Uploaded successfully")
}
