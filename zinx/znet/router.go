package znet

import "zinx/ziface"


/**
	在实现router时候，先定义一个BaseRouter基类，然后根据需要在创建别的路由并继承这个基类，然后在重写这个基类的方法即可。
	下面之所以BaseRouter的方法都为空，是因为有的Router不需要有PrevHandler()和PostHandler()这两个业务,为空的话，既满足了接口的实现规范，也不影响其他子类Router的方法定义
	所以，之后创建的Router继承了BaseRouter后，就必须要在实现未被实现的方法
 */
type BaseRouter struct {

}

//处理实际业务之前的方法
func (this *BaseRouter) PrevHandler(request ziface.IRequest){

}
//处理实际业务的主方法
func (this *BaseRouter) Handler(request ziface.IRequest){

}
//处理实际业务之后的方法
func (this *BaseRouter) PostHandler(request ziface.IRequest){

}