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

//TestPreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping....\n"))
	if err != nil {
		fmt.Println("callback before ping error", err)
	}
}

//TestHandle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping....ping....ping....\n"))
	if err != nil {
		fmt.Println("callback ping error", err)
	}
}

//TestPostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping....\n"))
	if err != nil {
		fmt.Println("callback after ping error", err)
	}
}

func main() {
	//1 创建一个server句柄，使用Zinx的APi
	s := znet.NewServer("[zinx V0.2]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
