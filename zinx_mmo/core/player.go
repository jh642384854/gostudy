package core

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
	"zinx/ziface"
)

/**
	定义玩家对象信息
 */
type Player struct {
	Pid int32  //玩家ID
	Conn ziface.IConnection  //当前玩家的链接(用于和客户端的链接)
	X float32  //平面的X轴坐标
	Y float32  //玩家跳跃的高度
	Z float32  //平面的Y轴坐标
	V float32  //玩家身体旋转的0-360角度
}

/**
	PlayerID生成器
 */
var PidGen int32 = 1  //用于生成玩家的ID计数器
var IdLock sync.Mutex //用来保护PidGen的锁

func NewPlayer(connection ziface.IConnection) *Player  {
	//生成一个玩家ID
	IdLock.Lock()
	pid := PidGen
	PidGen ++
	IdLock.Unlock()

	//创建一个玩家对象
	player := &Player{
		Pid:pid,
		Conn:connection,
		X:float32(160+rand.Intn(10)),//随机在160坐标点 基于X轴若干偏移
		Y:0,
		Z:float32(180+rand.Intn(20)),//随机在180坐标点，基于Y轴若干偏移
		V:0,
	}
	return player
}
/**
	提供一个发生客户端消息的方法
	主要是将pb的protobuf数据序列化之后，在调用zinx框架的SendMsg()方法
 */
func (this *Player) SendMsg(msgId uint32,data proto.Message)  {
	//判断链接是否为空，就是判断当前玩家是否处于在线状态
	if this.Conn == nil{
		fmt.Println("palyer outline")
		return
	}
	msg,err := proto.Marshal(data)
	if err != nil{
		fmt.Println("data Marsha1 error")
		return
	}
	if err := this.Conn.SendData(msgId,msg); err != nil{
		fmt.Println("player send message error")
		return
	}
	return
}