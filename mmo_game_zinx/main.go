package main

import (
	"fmt"
	"github.com/dongmyd/Code_Zinx/mmo_game_zinx/apis"
	"github.com/dongmyd/Code_Zinx/mmo_game_zinx/core"
	"github.com/dongmyd/Code_Zinx/zinx/ziface"
	"github.com/dongmyd/Code_Zinx/zinx/znet"
)

//当前客户端建立连接之后的hook函数
func OnConnectionAdd(conn ziface.IConnection) {
	//创建一个Player 的对象
	player := core.NewPlayer(conn)

	//给客户端发送MsgID：1的消息;同步当前Player的ID给客户端
	player.SyncPid()

	//给客户的发送MsgID：200的消息;同步当前Player的初始位置给客户端
	player.BroadCastStartPosition()

	//将当前新上线的玩家添加到WorldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该链接绑定一个Pid 玩家ID的属性
	conn.SetProperty("pid", player.Pid)

	// 同步周边玩家，告知他们当前玩家已经上线，广播当前玩家的位置信息
	player.SyncSurrounding()

	fmt.Println("========> Player pid = ", player.Pid, " is arrived <=====")
}

//给当前连接断开之前触发的Hook钩子函数
func OnConnectionLost(conn ziface.IConnection) {
	//通过连接属性得到当前连接所绑定pid
	pid, _ := conn.GetProperty("pid")

	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//触发玩家下线的业务
	player.Offline()

	fmt.Println("==========> Player pid = ", pid, " offline ... <===")
	//得到当前玩家周边的九宫格内都有哪些玩家
	//给周围玩家广播MsgID:201 消息
}

func main() {
	//创建zinx server句柄
	s := znet.NewServer("MMO Game Zinx")

	//连接创建和销毁的HOOK钩子函数
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	//注册一些路由业务
	s.AddRouter(2, &apis.WorldChatApi{})
	s.AddRouter(3, &apis.MoveApi{})

	//启动服务
	s.Server()
}
