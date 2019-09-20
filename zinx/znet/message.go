package znet

type Message struct {
	ID      uint32 //消息ID
	DataLen uint32 //消息长度
	Data    []byte //消息内容
}

//创建一个消息的方法
func NewMessage(mid uint32, data []byte) *Message {
	return &Message{
		ID:      mid,
		DataLen: uint32(len(data)), //注意这里要做一个uint32的强转
		Data:    data,
	}
}

//获取消息ID
func (this *Message) GetMessageID() uint32 {
	return this.ID
}

//获取消息长度
func (this *Message) GetMessageLen() uint32 {
	return this.DataLen
}

//获取消息内容
func (this *Message) GetMessageData() []byte {
	return this.Data
}

//设置消息ID
func (this *Message) SetMessageID(id uint32) {
	this.ID = id
}

//设置消息长度
func (this *Message) SetMessageLen(len uint32) {
	this.DataLen = len
}

//设置消息内容
func (this *Message) SetMessageData(data []byte) {
	this.Data = data
}
