package znet

import "github.com/dongmyd/Code_Zinx/zinx/ziface"

//实现router时，先嵌入BaseRouter积累，然后根据需要对这个基类的方法进行重写就好
type BaseRouter struct {
}

//这里之所以BaseRouter的方法都为空
//是因为有的Router不希望有PerHandle,PostHandle这两个业务
//所以Router全部继承BaseRouter的好处就是，不需要实现
// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PerHandle(requset ziface.IRequest) {}

//在处理conn业务的主方法Hook
func (br *BaseRouter) Handle(requset ziface.IRequest) {}

//在处理conn业务之后的钩子方法Hook
func (br *BaseRouter) PostHandle(requset ziface.IRequest) {}
