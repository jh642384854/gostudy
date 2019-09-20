package pool

import (
	"github.com/pkg/errors"
	"io"
	"log"
	"sync"
)

/**
	这个包用于展示如何使用有缓冲的通道实现资源池，来管理可以在任意数量的goroutine之间共享及独立使用的资源。
	这种模式在需要共享一组静态资源的情况（如共享数据库连接或者内存缓冲区）下非常有用。
	如果goroutine需要从池里得到这些资源中的一个，它可以从池里申请，使用完后归还到资源池里

	在 Go 1.6 及之后的版本中，标准库里自带了资源池的实现（sync.Pool）,推荐使用。
 */

//Pool管理一组可以安全地在多个goroutine之间共享资源，被管理者(这些资源)必须实现io.Closer接口
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed表示请求(Acquire)了一个已经关闭的池
var ErrPoolClosed = errors.New("Pool has been closed")

// New创建一个用来管理资源的池。这个池需要一个可以分配新资源的函数，并规定池的大小
func New(fn func()(io.Closer,error),size uint) (*Pool,error)  {
	if size <= 0{
		return nil,errors.New("Size Value To Small.")
	}
	return &Pool{
		factory:fn,
		resources:make(chan io.Closer,size),
	},nil
}

//Acquire 从池中获取一个资源
func (p *Pool) Acquire() (io.Closer,error)  {
	select {
		case r,ok := <- p.resources:
			log.Println("Acquire:shared resource")
			if !ok{
				return nil,ErrPoolClosed
			}
			return r,nil
	default:
		log.Println("Acquire：New Resource")
		return p.factory()
	}
}

//Release 将一个使用后的资源放回池里
func (p *Pool) Release(r io.Closer)  {
	//保证本操作和Close操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果池已经被关闭，销毁这个资源
	if p.closed{
		r.Close()
		return
	}

	select {
	//试图将这个资源放入队列
	case p.resources <- r:
		log.Println("Release:In Queue")
	//如果队列已满，则关闭这个资源
	default:
		log.Println("Release:Closing")
		r.Close()
	}
}
// Close会让资源池停止工作，并关闭所有现有的资源。
func (p *Pool) Close()  {
	//保证本操作与Release操作的安全
	p.m.Lock()
	defer p.m.Unlock()

	//如果pool已经被关闭，什么都不做
	if p.closed{
		return
	}

	//将池关闭
	p.closed = true

	//在清空通道里的资源之前，将通道关闭，如果不这样做，会发生死锁
	close(p.resources)

	//关闭资源
	for r := range p.resources {
		r.Close()
	}
}