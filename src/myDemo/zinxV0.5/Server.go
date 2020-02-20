//Time    : 2020-02-18 17:14
//Author  : Hugh
//File    : Server.go
//Descripe:

package main

import (
	"awesomeProject/src/zinx/ziface"
	"awesomeProject/src/zinx/znet"
	"fmt"
)

/*
	基于Zinx框架来开发的服务器端应用程序
*/
type PingRouter struct {
	znet.BaseRouter
}

//TestHandle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client :msgID=", request.GetMsgId(), " data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping....ping...ping...."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//1 创建一个server句柄，使用Zinx的APi
	s := znet.NewServer("[zinx V0.5]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
