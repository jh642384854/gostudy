package znet

import (
	"errors"
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

/**
	多消息处理模块
 */

type MsgHandler struct {
	Apis         map[uint32]ziface.IRouter
	WorkPoolSize uint32                 //定义工作池大小
	TaskQueue    []chan ziface.IRequest //定义每个工作池中任务队列中数据大小
}

//创建MsgHandler的实例
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:         make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GloabConfigObj.MaxWorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GloabConfigObj.MaxWorkPoolSize),
	}
}

//执行指定消息的路由信息
func (this *MsgHandler) DoMsgHandler(request ziface.IRequest) error {
	//判断要被调用的消息id对应的router是否存在
	handler, ok := this.Apis[request.GetDataID()]
	if !ok {
		fmt.Println("该router没有被被注册，无法使用该功能")
		return errors.New("该router已经被注册，无需重复注册")
	}
	//取出是那个router处理器，然后依次调用前置、后置和实际处理器
	handler.PrevHandler(request)
	handler.Handler(request)
	handler.PostHandler(request)
	return nil
}

//添加路由到消息管理器中
func (this *MsgHandler) AddMsgHandler(msgid uint32, router ziface.IRouter) error {
	//判断是否已经注册过了
	if _, ok := this.Apis[msgid]; ok {
		fmt.Println("该router已经被注册，无需重复注册")
		return errors.New("该router已经被注册，无需重复注册")
	}
	this.Apis[msgid] = router
	return nil
}

//启动工作池
func (this *MsgHandler) StartWorkPool() {
	//根据设置的最大工作池数量来开启响应的go routine来处理
	for i := 0; i < int(utils.GloabConfigObj.MaxWorkPoolSize); i++ {
		//为每个工作池单独分配空间
		this.TaskQueue[i] = make(chan ziface.IRequest, utils.GloabConfigObj.MaxTaskQueueLen)
		go this.startOneWork(i, this.TaskQueue[i])
	}
}

//开启每个工作池的任务
func (this *MsgHandler) startOneWork(workid int, workRequest chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workid, " is started ...")
	//循环阻塞处理用户请求
	for {
		select {
		case request := <-workRequest:
			this.DoMsgHandler(request)
		}
	}
}

//将处理请求加入到工作池中
func (this *MsgHandler) AddRequestToTaskQueue(workRequest ziface.IRequest) {
	//一个简单的算法，将当前的请求存放到哪个TaskQueue中
	workid := workRequest.GetConnect().GetConnectionID() & utils.GloabConfigObj.MaxWorkPoolSize
	this.TaskQueue[workid] <- workRequest
}
