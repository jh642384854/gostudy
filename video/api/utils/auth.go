package utils

import (
	"dev/video/api/defs"
	"dev/video/api/session"
	"net/http"
)

var (
	HEADER_FIELD_SESSION = "X-Session-Id"
	HEADER_FIELD_UNAME = "X-User-Name"
)
//校验用户session
func ValidateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0{
		return  false
	}
	if username,ok := session.IsSessionExpired(sid); ok{
		r.Header.Add(HEADER_FIELD_UNAME,username)
		return true
	}else{
		return false
	}
}

//校验用户信息
func ValidateUser(w http.ResponseWriter,r *http.Request) bool {
	username := r.Header.Get(HEADER_FIELD_UNAME)
	if len(username) == 0{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}