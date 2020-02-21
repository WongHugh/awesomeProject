//Time    : 2020-02-18 16:54
//Author  : Hugh
//File    : iserver.go
//descripe:

package ziface

//定义一个服务器接口
type IServer interface {
	Start()
	Stop()
	Serve()

	//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
	AddRouter(msgId uint32, router IRouter)
	//获取当前server的连接管理器
	GetConnMgr() IConnManager
	//注册OnConnStart钩子函数的方法
	SetOnConnStart(func(connection IConnection))
	//注册OnConnStop钩子函数的方法
	SetOnConnStop(func(connection IConnection))
	//调用OnConnStart钩子函数的方法
	CallOnConnStart(connection IConnection)
	//调用OnConnStop钩子函数的方法
	CallOnConnStop(connection IConnection)
}
