//Time    : 2020-02-18 16:55
//Author  : Hugh
//File    : server.go
//descripe:

package znet

import (
	"awesomeProject/src/zinx/utils"
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
	//当前的Server 的消息管理模块，用来绑定msgID和对应的处理业务的API关系
	MsgHandler ziface.IMsgHandle
	//该server的链接管理器
	ConnMgr ziface.IConnManager

	OnConnStart func(conn ziface.IConnection)
	OnConnStop  func(conn ziface.IConnection)
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
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
	fmt.Printf("[Zinx] serverName :%s,listenner at IP:%s,Port:%d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s,MaxConn:%d,MaxPackageSize:%d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start] Server Listenner at IP :%s, Port:%d is starting\n", s.IP, s.Port)

	go func() {
		s.MsgHandler.StartWorkerPoll()
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
			//最大链接个数判断
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//todo 回复客户端已超出最大数量
				fmt.Println("too many connections,MaxConn=", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)

			cid++
			go dealConn.Start()
		}

	}()

}
func (s *Server) Stop() {
	fmt.Println("[Stop] Zinx server,name", s.Name)
	//TODO 将一些服务器的资源、状态或一些开辟的链接信息进行停止或者回收
	s.ConnMgr.ClearConn()

}
func (s *Server) Serve() {
	s.Start()

	//TODO 做一些启动服务器之后的业务
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

/*
	初始化Server模块的方法
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}

func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

//注册OnConnStop钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("--->Call OnConnStart....")
		s.OnConnStart(conn)
	}
}

//调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("--->Call OnConnStop....")
		s.OnConnStop(conn)
	}
}
