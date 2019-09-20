package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

/**
	封包、拆包模块
	直接面向TCP连接中的数据流，用于处理TCP粘包问题
 */

type DataPack struct {

}

//拆包封包的一个初始化方法
func NewDataPack() *DataPack  {
	return  &DataPack{}
}

func (this *DataPack) GetHeadLen() uint32{
	/**
		这里是我们协议约定的。下面的数字8就是由datalen和msgid组成，这两个各占4个字节，所以这里就是8
		一个消息体的构成：
		datalen + msgid + data 【datalen是消息长度，msgid是消息id，data消息的实际内容】
		其中datalen + msgid 就是消息head(消息头部)，而data就是body(消息实际内容)
		为什么会有一个消息ID呢？
		我们可以自己来设定一些消息ID，根据这些不同的消息ID来处理不同的消息内容。
	 */
	return  8
}

//封包的方法。将Message对象转换为二进制格式，进行网络传输
func (this *DataPack) Pack(msg ziface.IMessage) ([]byte,error){
	dataBuff := bytes.NewBuffer([]byte{})

	//将datalen写进databuff中
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageLen());err != nil{
		return  nil,err
	}
	//将MsgID写进databuff中
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageID());err != nil{
		return  nil,err
	}
	//将data写进databuff中
	if err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMessageData());err != nil{
		return  nil,err
	}
	//将二进制序列化返回
	return dataBuff.Bytes(),nil
}

//拆包的方法。将读取的网络二进制数据转换为我们定义的Message对象。
//将包的head信息读取出来，之后在根据head信息里的data长度，再进行一次读
func (this *DataPack) Unpack(binaryData []byte) (ziface.IMessage,error){
	//创建一个从输入二进制的ioReader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head信息，得到datalen和MsgID
	msg := &Message{}

	//依次从二进制流中读取之前封装好的信息
	//首先读取datalen
	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.DataLen); err != nil{
		return  nil,err
	}
	//然后读取MsgID
	if err := binary.Read(dataBuff,binary.LittleEndian,&msg.ID); err != nil{
		return  nil,err
	}
	//判断读取的数据量是否大于我们在src/zinx/utils/config.go里面中GloabConfig定义的MaxPackageSize这个值
	if (utils.GloabConfigObj.MaxPackageSize > 0 && msg.DataLen > utils.GloabConfigObj.MaxPackageSize){
		return  nil,errors.New("发送的数据流过大")
	}
	//此时返回的msg对象是没有实际的数据内容的，只有datalen和msgid。但是我们可以根据msg对象的datalen属性再次从链接中读取信息，从而获取数据的的实际内容
	return msg,nil
}