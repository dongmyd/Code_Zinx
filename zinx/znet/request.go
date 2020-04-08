package znet

import "github.com/dongmyd/Code_Zinx/zinx/ziface"

type Requset struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection

	//客户端请求的数据
	data []byte
}

func (r *Requset) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Requset) GetData() []byte {
	return r.data
}
