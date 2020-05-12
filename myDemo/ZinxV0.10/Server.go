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

//创建链接之后执行钩子函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionBegin is Called ...")
	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}

	//给当前的链接设置一些属性
	fmt.Println("Set conn Name, Hoe ...")
	conn.SetProperty("Name", "刘丹冰-Aceld-IT无崖子")
	conn.SetProperty("GitHub", "https://github.com/aceld")
	conn.SetProperty("Home", "https://legacy.gitbook.com/@aceld")
	conn.SetProperty("Blog", "https://www.jianshu.com/u/35261429b7f1")
}

//链接断开之前的需要执行的函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionLost is Called ... ")
	fmt.Println("conn ID = ", conn.GetConnID(), " is Lost...")

	//获取链接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Home = ", home)
	}
	if gitHub, err := conn.GetProperty("GitHub"); err == nil {
		fmt.Println("GitHub = ", gitHub)
	}
	if blog, err := conn.GetProperty("Blog"); err == nil {
		fmt.Println("Blog = ", blog)
	}
}

func main() {
	//1.创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx V0.9]")

	//2 注册链接Hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//3.给当前zinx框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	//4.启动server
	s.Server()

}
