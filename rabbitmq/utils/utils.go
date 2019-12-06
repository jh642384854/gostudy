package utils

import (
	"log"
	"os"
	"strings"
)

func FailOnError(err error,msg string)  {
	if err != nil{
		log.Fatalf("%s:%s",err,msg)
	}
}
/**
	自己传递消息内容
 */
func BodyForm(args []string) string {
	var s string
	if (len(args) <2 || os.Args[1] == ""){
		s = "hello"
	}else{
		s = strings.Join(args[1:]," ")
	}
	return s
}

/**
	发送消息的时候指定消息类型，这个用来设置routing_key
 */
func SeverityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}