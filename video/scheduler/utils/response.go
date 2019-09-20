package utils

import (
	"io"
	"net/http"
)

func SendMessage(w http.ResponseWriter,statusCode int,msg string)  {
	w.WriteHeader(statusCode)
	io.WriteString(w,msg)
}