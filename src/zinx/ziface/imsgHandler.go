//Time    : 2020-02-20 14:48
//Author  : Hugh
//File    : imsgHandler.go
//Descripe:

package ziface

/*
	消息管理抽象层
*/
type IMsgHandle interface {
	//调度执行对应的router消息处理
	DoMsgHandle(request IRequest)
	//为消息添加具体的处理逻辑
	AddRouter(msgId uint32, router IRouter)
}
