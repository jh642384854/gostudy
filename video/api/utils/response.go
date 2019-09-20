package utils

import (
	"dev/video/api/defs"
	"encoding/json"
	"io"
	"net/http"
)

//错误信息输出
func sendErrorResponse(w http.ResponseWriter,errRes defs.ErrorResponse)  {
	w.WriteHeader(errRes.HttpStatusCode)
	resStr,_ := json.Marshal(errRes.Error)
	io.WriteString(w,string(resStr))
}

//正常信息输出
func sendNormalResponse(w http.ResponseWriter,message string,code int)  {
	w.WriteHeader(code)
	io.WriteString(w,message)
}