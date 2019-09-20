package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//只是负责测试databack拆包、封包的单元测试
func TestDataPack(t *testing.T) {

	//1.模拟服务器
	//①、建立TCP链接
	listener,err := net.Listen("tcp","127.0.0.1:7777")
	if err != nil{
		fmt.Println("server listener err",err)
		return
	}

	//②、创建一个go routine 用来承载负责从客户端发送过来的请求业务
	go func() {
		//③、从客户端读取数据，进行拆包处理

		//循环读取客户端发送的消息
		for {
			//阻塞
			conn,err := listener.Accept()
			if err != nil{
				fmt.Println("server Accept err",err)
			}
			//实际处理客户端请求。这里实际就是一个拆包的过程
			go func(conn net.Conn) {
				//定义一个拆包对象
				dp := NewDataPack()
				for  {
					//第一次从conn中读取数据，把包中的head读取出来
					//定义一个byte切片，将从conn读取head数据。
					headdata := make([]byte,dp.GetHeadLen())
					_,err := io.ReadFull(conn,headdata)
					if err != nil{
						fmt.Println("server Read err",err)
						break
					}
					//将读取的二进制流进行解包操作，得到的就是只有包含消息head的Message对象
					messageHead,err := dp.Unpack(headdata)
					if err != nil{
						fmt.Println("server UnpackData err",err)
						return
					}
					//判断消息的长度是否真正大于0,表示有数据长度
					// 第二次从conn中读取数据，就会根据head中的datalen在读取实际的内容
					if messageHead.GetMessageLen() >0 {
						msg := messageHead.(*Message)
						msg.Data = make([]byte,msg.GetMessageLen())
						//根据datalen的长度再次从io流中读取
						_,err := io.ReadFull(conn,msg.Data)
						if err != nil{
							fmt.Println("server readData err",err)
							return
						}
						//到此消息读取完毕，可以显示消息
						fmt.Printf("Msgid:%d，data len :%d，body :%s \n",msg.ID,msg.DataLen,string(msg.Data))
					}
				}
			}(conn)
		}
	}()

	//2.模拟客户端
	//①、客户端建立链接
	conn,err := net.Dial("tcp","127.0.0.1:7777")
	if err != nil{
		fmt.Println("client contact err",err)
		return
	}
	//②、创建一个封包对象，将要发送的消息进行封包处理
	datapack := NewDataPack()

	//③、定义几条消息，然后一起发送到服务端
	msg1 := &Message{
		ID:1,
		DataLen:5,
		Data:[]byte{'h','e','l','l','o'},
	}
	msgByte1,err := datapack.Pack(msg1)
	if err != nil{
		fmt.Println("datapack pack err",err)
		return
	}

	msg2 := &Message{
		ID:2,
		DataLen:4,
		Data:[]byte{'z','i','n','x'},
	}
	msgByte2,err := datapack.Pack(msg2)
	if err != nil{
		fmt.Println("datapack pack err",err)
		return
	}
	//将以上两个消息粘在一起
	msgByte2 = append(msgByte2,msgByte1...)
	conn.Write(msgByte2)
	//客户端阻塞
	select {

	}
}