package main

import (
	"dev/day5/pool"
	"time"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
)

const (
	maxGoroutines = 25 //要使用的goroutine数量
	pooledResoureces = 2 //池中资源的数量
)
//dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

//Close实现了io.Close接口，以便dbConnection可以被池管理。Close()用来完成任意这样的师傅
func (dbConn *dbConnection) Close() error  {
	log.Println("Close:Connection",dbConn.ID)
	return nil
}
//idCounter用来给每一个连接分配一个独一无二的ID
var idCounter int32

//这个是一个工厂函数，当需要一个新连接时候，资源池会调用这个函数
func createConnection() (io.Closer,error) {
	id := atomic.AddInt32(&idCounter,1)
	log.Println("Create:New Connection", id)
	return &dbConnection{id},nil

}

func main() {
	var wg8 sync.WaitGroup
	wg8.Add(maxGoroutines)

	p,err := pool.New(createConnection,pooledResoureces)
	if err != nil{
		log.Println(err)
	}
	for query := 0;query < maxGoroutines ; query++ {
		//每个goroutine需要自己复制一份要查询值的副本，不然所有的查询会共享同一个查询变量
		go func(q int) {
			performQueries(q,p)
			wg8.Done()
		}(query)
	}
	//等待goroutine结束
	wg8.Wait()
	//关闭池
	log.Println("shutdown program")
	p.Close()
}
// 这个用来测试连接的资源池
func performQueries(query int,p *pool.Pool)  {
	//从池里面请求一个连接
	conn,err := p.Acquire()
	if err != nil{
		log.Println(err)
		return
	}
	//将该连接释放会池里
	defer p.Release(conn)
	//用等待来模拟查询响应
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	log.Printf("QID[%d] CID[%d] \n",query,conn.(*dbConnection).ID)
}

