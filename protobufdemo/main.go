package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"protobufdemo/pb"
)

func main() {
	//依据Person.proto这个消息格式文件来创建一下消息内容(当然这些消息的创建最终是依据Person.proto这个生成的Person.pb.go文件格式来创建的)。
	person := &pb.Person{
		Id:1,
		Name:"zhangsan",
		Phones:[]*pb.Phone{
			&pb.Phone{
				Number:"15124251242",
				Type:pb.PhoneType_HOME,
			},
			&pb.Phone{
				Number:"1570271963322",
				Type:pb.PhoneType_WORK,
			},
			&pb.Phone{
				Number:"157071896332",
				Type:pb.PhoneType_HOME,
			},
		},
	}
	//使用protobuf来编码消息
	data,err := proto.Marshal(person)
	if err != nil{
		fmt.Println("protobuf Marsha1 fail")
	}
	//解码信息
	newPerson := &pb.Person{}
	if err := proto.Unmarshal(data,newPerson); err != nil{
		fmt.Println("protobuf Unmarshal fail")
	}
	fmt.Println("解码前数据：",person)
	fmt.Println("解码后数据：",newPerson)
}
