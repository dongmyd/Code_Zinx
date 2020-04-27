package main

import (
	"fmt"
	"github.com/dongmyd/Code_Zinx/zinx/znet"
	"io"
	"net"
	"time"
)

/*模拟客户端*/
func main() {
	fmt.Println("client0 start ...")

	time.Sleep(1 * time.Second) //睡一秒方便显示查看效果

	//1 连接远程服务器，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start er, exit!")
		return
	}

	for {
		//发送封包message消息 MsgID:0
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx client0 Test Message")))
		if err != nil {
			fmt.Println("Pack error:", err)
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error :", err)
			return
		}

		//服务器就应该给我们回复一个message数据，MsgID:pingpingping

		//1 先读取流中的head部分 得到ID 和 dataLen

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error ", err)
			break
		}
		//将二进制的head拆包到msg 结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//2再跟进DataLen进行第二次读取，将data读出来
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data error,", err)
				return
			}

			fmt.Println("----> Recv Server Msg:ID = ", msg.Id, ", len = ", msg.DataLen, ", data = ", string(msg.Data))
		}

		//阻塞 预防cpu跑满 无限循环 每隔一秒write一次给服务器
		time.Sleep(1 * time.Second)
	}
}
