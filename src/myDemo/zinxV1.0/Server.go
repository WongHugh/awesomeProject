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
	err := request.GetConnection().SendMsg(200, []byte("ping....ping...ping...."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

//TestHandle
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle...")
	fmt.Println("recv from client :msgID=", request.GetMsgId(), " data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(201, []byte("Hello welcome to zinx"))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is called...")
	if err := conn.SendMsg(202, []byte("DoConnectionBegin...")); err != nil {
		fmt.Println(err)
	}

	//给当前的连接设置一些属性
	fmt.Println("Set conn Name,Hoe....")
	conn.SetProperty("Name", "Hugh")
	conn.SetProperty("Github", "http://bi.cn")

}

func DoConnectionAfter(conn ziface.IConnection) {
	fmt.Println("=======>DoConnectionAfter is called...")
	fmt.Println("connID =", conn.GetConnID(), " is closed")

	//获取链接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name", name)
	}

	if Github, err := conn.GetProperty("Github"); err == nil {
		fmt.Println("Name", Github)
	}
}

func main() {
	//1 创建一个server句柄，使用Zinx的APi
	s := znet.NewServer("[zinx V1.0.0]")
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionAfter)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()
}
