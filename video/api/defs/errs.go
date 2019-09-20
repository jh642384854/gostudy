package defs

//定义错误信息
type Error struct {
	ErrorMsg string `json:"error_msg"`
	ErrorCode string `json:"error_code"`  //错误状态码
}
//定义HTTP请求响应的错误信息
type ErrorResponse struct {
	HttpStatusCode int `json:"http_status_code"`
	Error Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpStatusCode:400,
		Error:Error{
			ErrorMsg:"Request body is not correct",
			ErrorCode:"001",
		},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpStatusCode:401,
		Error:Error{
			ErrorMsg:"User Authentication failed",
			ErrorCode:"002",
		},
	}
	ErrorDbOperate = ErrorResponse{
		HttpStatusCode:500,
		Error:Error{
			ErrorMsg:"Database operation failed",
			ErrorCode:"003",
		},
	}
	ErrorInternalFaults = ErrorResponse{
		HttpStatusCode:500,
		Error:Error{
			ErrorMsg:"Server Internal Error",
			ErrorCode:"004",
		},
	}
)