//Time    : 2020-02-18 17:14
//Author  : Hugh
//File    : Server.go
//Descripe:

package main

import "awesomeProject/src/zinx/znet"

/*
	基于Zinx框架来开发的服务器端应用程序
*/
func main() {
	//1 创建一个server句柄，使用Zinx的APi
	s := znet.NewServer("[zinx V0.2]")
	s.Serve()
}
