package utils

import (
	"dev/video/api/dbops"
	"dev/video/api/defs"
	"dev/video/api/session"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	//解析提交的用户数据
	res,_ := ioutil.ReadAll(r.Body)
	fmt.Println(string(res))
	userObj := &defs.UserCredential{}
	if err := json.Unmarshal(res,userObj);err != nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	//在用户表中添加一条记录
	if err := dbops.AddUserCredential(userObj.UserName,userObj.Pwd);err != nil{
		sendErrorResponse(w,defs.ErrorDbOperate)
		return
	}
	//在session表中添加一条记录
	session_id := session.GenerateNewSessionID(userObj.UserName)
	singedUp := &defs.SignedUp{Success:true,SessionId:session_id}
	//将成功添加的信息进行返回转成JSON后输出
	if res,err := json.Marshal(singedUp); err != nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(w,string(res),201)
	}
}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	io.WriteString(w,p.ByName("user_name"))
}