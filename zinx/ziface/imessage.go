package ziface

/**
	消息内容的封装
 */

type IMessage interface {
	//获取消息ID
	GetMessageID() uint32
	//获取消息长度
	GetMessageLen() uint32
	//获取消息内容
	GetMessageData() []byte

	//设置消息ID
	SetMessageID(id uint32)
	//设置消息长度
	SetMessageLen(len uint32)
	//设置消息内容
	SetMessageData(data []byte)

}
