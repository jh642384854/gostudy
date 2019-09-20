package znet

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"sync"
	"zinx/ziface"
)

type ConnectionManager struct {
	//记录链接信息，使用map数据结构，下标就是链接ID，值就是链接对象
	Connections map[uint32]ziface.IConnection
	ConnLock sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		Connections:make(map[uint32]ziface.IConnection),
	}
}

//添加链接
func (this *ConnectionManager) Add(conn ziface.IConnection){
	//保护共享资源map， 加写锁
	this.ConnLock.Lock()
	defer this.ConnLock.Unlock()
	this.Connections[conn.GetConnectionID()] = conn
}
//根据链接ID来删除链接
func (this *ConnectionManager) Remove(connid uint32){
	//保护共享资源map， 加写锁
	this.ConnLock.Lock()
	defer this.ConnLock.Unlock()
	//判断是否存在
	if _,ok := this.Connections[connid];ok{
		delete(this.Connections,connid)
	}else{
		fmt.Println("当前链接不存在，无法删除")
	}
}
//根据链接ID来获取链接
func (this *ConnectionManager) Get(connid uint32) (ziface.IConnection,error){
	//保护共享资源map， 加读锁
	this.ConnLock.RLocker()
	defer this.ConnLock.RUnlock()
	if conn,ok := this.Connections[connid];ok{
		return conn,nil
	}else{
		fmt.Println("当前链接不存在，无法查找")
		return nil,errors.New("当前链接不存在，无法查找")
	}
}
//得到当前链接管理器的所有链接总数。这个用来控制流量，设置当前允许最大链接数
func (this *ConnectionManager) Len() int{
	return len(this.Connections)
}
//清除所有链接
func (this *ConnectionManager) ClearAll(){
	//保护共享资源map， 加写锁
	this.ConnLock.Lock()
	defer this.ConnLock.Unlock()
	for connid,conn := range this.Connections {
		//停止服务。这个很重要
		conn.Stop()
		//删除链接
		delete(this.Connections,connid)
	}
}