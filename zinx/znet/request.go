package znet

import (
	"zinx/ziface"
)

type Request struct {
	Conn ziface.IConnection
	//Data []byte
	msgData ziface.IMessage
}

func (this *Request) GetConnect() ziface.IConnection {
	return this.Conn
}

//得到当前的消息数据
func (this *Request) GetData() []byte {
	//return this.Data
	return this.msgData.GetMessageData()
}

//得到当前消息的ID
func (this *Request) GetDataID() uint32 {
	return this.msgData.GetMessageID()
}
