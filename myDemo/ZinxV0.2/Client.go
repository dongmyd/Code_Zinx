package main

import (
	"fmt"
	"net"
	"time"
)

/*模拟客户端*/
func main()  {
	fmt.Println("client start ...")

	time.Sleep(1*time.Second)//睡一秒方便显示查看效果

	//1 连接远程服务器，得到一个conn链接
	conn,err := net.Dial("tcp","127.0.0.1:8999")
	if err != nil{
		fmt.Println("client start er, exit!")
		return
	}

	for  {
		//2 链接调用Write 写数据
		_,err:= conn.Write([]byte("Hello Zinx V0.2..."))
		if err!=nil{
			fmt.Println("write conn err",err)
			return
		}

		buf := make([]byte,512) //用buf接收服务器返回的数据
		cnt, err := conn.Read(buf) //读取服务器返回的数据
		if err !=nil{
			fmt.Println("read buf error")
			return
		}

		fmt.Printf(" server call back :buf = %s, cnt = %d\n",buf,cnt)

		//阻塞 预防cpu跑满 无限循环 没隔一秒write一次给服务器
		time.Sleep(1*time.Second)
	}
}