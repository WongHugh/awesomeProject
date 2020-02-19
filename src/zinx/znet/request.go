//Time    : 2020-02-19 14:11
//Author  : Hugh
//File    : request.go
//Descripe:

package znet

import "awesomeProject/src/zinx/ziface"

type Request struct {
	//已经和客户端建立好的链接
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

//得到链接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.data
}
