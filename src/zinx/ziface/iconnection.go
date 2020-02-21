//Time    : 2020-02-19 11:20
//Author  : Hugh
//File    : iconnection.go
//Descripe:

package ziface

import "net"

// 定义链接模块的抽象层
type IConnection interface {
	//启动链接，让当前的链接准备工作
	Start()
	//停止链接，结束当前链接的工作
	Stop()
	//获取当前链接的绑定sokect conn
	GetTCPConnection() *net.TCPConn
	//获取当前链接模块的链接ID
	GetConnID() uint32
	//获取远程客户端的TCP状态 IP Port
	RemoterAddr() net.Addr
	//发送数据，将舒服发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error

	//设置链接属性
	SetProperty(key string, value interface{})
	//获取链接属性
	GetProperty(key string) (interface{}, error)

	//删除链接熟悉
	RemoveProperty(key string)
}

//定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
