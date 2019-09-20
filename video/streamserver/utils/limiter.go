package utils

import "log"

/**
	流量控制
 */

type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(count int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn:count,
		bucket:make(chan int,count),
	}
}

//获得一个连接
func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.concurrentConn{
		log.Printf("Reached the rate limitation.")
		return  false
	}
	cl.bucket <- 1
	return true
}

//释放一个连接
func (cl *ConnLimiter) ReleaseConn()  {
	c := <- cl.bucket
	log.Printf("New connection coming : %d",c)
}

