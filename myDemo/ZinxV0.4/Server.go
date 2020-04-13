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

//Test PerHandle
func (this *PingRouter) PerHandle(request ziface.IRequest) {
	fmt.Println("Call Router PerHandle")
	_, err := request.GetConnection().GetTCPCOnnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := request.GetConnection().GetTCPCOnnection().Write([]byte("ping...ping...ping...\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping.. error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPCOnnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	//1.创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.4]")

	//2.给当前zinx框架添加一个自定义的router
	s.AddRouter(&PingRouter{})

	//2.启动server
	s.Server()

}
