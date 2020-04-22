package main

import (
	"fmt"
	"github.com/dongmyd/Code_Zinx/zinx/ziface"
	"github.com/dongmyd/Code_Zinx/zinx/znet"
)

/*
	基于Zinx框架来开发的 服务端应用程序
*/

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping..ping..ping
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		" , data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}

}

//HelloZinxRouter 自定义路由
type HelloZinxRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping..ping..ping
	fmt.Println("recv from client:msgID = ", request.GetMsgID(),
		" , data = ", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("Hello Welcome to Zinx!!"))
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	//1.创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.5]")

	//2.给当前zinx框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//2.启动server
	s.Server()

}
