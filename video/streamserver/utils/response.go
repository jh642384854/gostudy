package utils

import (
	"io"
	"net/http"
)

/**
	响应处理。
 */

func SendErrorResponse(w http.ResponseWriter,statucCode int,errMsg string)  {
	w.WriteHeader(statucCode)
	io.WriteString(w,errMsg)
}