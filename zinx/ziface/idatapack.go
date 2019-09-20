package ziface

/**
	数据包的解包和封包操作，解决TCP粘包问题
	采用的TLV(Type-Lenth-Value)格式来进行。下面简单的来说明一下什么是TLV格式：
	我们会采用下面的组合方式来定义一个数据包：
	消息类型+消息长度+消息实际内容
	其中消息类型和消息长度这两个内容我们称为head，并且这个head是一个固定的二进制长度。

 */

type IDatapack interface {
	//得到消息长度
	GetHeadLen() uint32
	/**
		针对Message进行TLV格式封装
		先写消息长度，在写消息ID，在写消息的实际内容
		将消息内容转换成二进制流格式
	 */
	Pack(msg IMessage) ([]byte,error)
	/**
		针对Message进行TLV格式拆包，这个需要分为两步来进行：
		①、先读取固定长度的head，通过固定长度的head来得到消息的真正长度和消息的类型。
		②、在根据消息内容的长度，再次进行下一次读取，从conn中读取实际消息的内容
		从二进制流中解析消息，并封装成IMessage对象
	 */

	Unpack([]byte) (IMessage,error)
}