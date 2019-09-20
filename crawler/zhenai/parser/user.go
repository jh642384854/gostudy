package parser


import (
	"dev/crawler/engine"
	"dev/crawler/model"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

//返回的结果是Request，这个对象里面包含了request和items
func ParseUser(content []byte) engine.ParseResult  {
	regExpression := `window.__INITIAL_STATE__=.*};`
	reg := regexp.MustCompile(regExpression)
	regResult := reg.Find(content)
	jsonData := strings.TrimRight(strings.TrimLeft(string(regResult),"window.__INITIAL_STATE__="),";")
	user,_ := decodeJson(jsonData)
	result := engine.ParseResult{}
	result.Items = append(result.Items,user)
	return result
}

func decodeJson(jsondata string) (model.User,error) {
	//定义一个通用的结构。
	var user model.User
	var userinfo map[string]interface{}
	//var userinfo = make(map[string]interface{})
	if err := json.Unmarshal([]byte(jsondata),&userinfo); err != nil{
		return user,fmt.Errorf("err : %v",err)
	}
	userJson,err := json.Marshal(userinfo["objectInfo"])
	if err !=nil{
		return user,fmt.Errorf("err : %v",err)
	}
	if err := json.Unmarshal(userJson,&user); err !=nil{
		return user,fmt.Errorf("err : %v",err)
	}
	return user,nil
}