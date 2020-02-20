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
}
