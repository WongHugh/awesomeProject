//Time    : 2020-02-18 16:55
//Author  : Hugh
//File    : server.go
//descripe:

package znet

import (
	"awesomeProject/src/zinx/ziface"
	"errors"
	"fmt"
	"net"
	"time"
)

//IServer 的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
	//当前的Server 添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success!")
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[ConnHandle]CallBackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP :%s, Port:%d is starting\n", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("Resolve tcp addr error :", err)
		}

		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, " err", err)
			return
		}
		fmt.Println("Start Zinx server success, ", s.Name, ", Listening...")
		var cid uint32
		cid = 0
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router)

			cid++
			go dealConn.Start()
		}

	}()

}
func (s *Server) Stop() {
	fmt.Println("[Stop] Zinx server,name", s.Name)
	//TODO 将一些服务器的资源、状态或一些开辟的链接信息进行停止或者回收

}
func (s *Server) Serve() {
	s.Start()

	//TODO 做一些启动服务器之后的业务
	for {
		time.Sleep(10 * time.Second)
	}
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      10010,
		Router:    nil}
	return s
}
