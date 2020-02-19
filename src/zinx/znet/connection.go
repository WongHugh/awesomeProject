//Time    : 2020-02-19 11:28
//Author  : Hugh
//File    : connection.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/ziface"
	"fmt"
	"net"
)

/*
	链接模块
*/
type Connection struct {
	//当前链接的socket TCP 套接字
	Conn *net.TCPConn
	//链接的ID
	ConnId uint32
	//当前的链接状态
	isClosed bool
	//当前链接所绑定的处理业务的方法API
	handleAPI ziface.HandleFunc
	// 告知当前链接已经退出/停止 channel
	ExitChal chan bool
}

//初始化链接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connID,
		isClosed:  false,
		handleAPI: callback_api,
		ExitChal:  make(chan bool, 1),
	}
	return c
}

//链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Read Gorouting is running...")
	defer fmt.Println("connID=", c.ConnId, "Reader is exit,remote addr is ", c.RemoterAddr().String())
	defer c.Stop()
	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buff err", err)
			c.ExitChal <- true
			continue
		}
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("connID=", c.ConnId, err)
			c.ExitChal <- true
			break
		}

	}

}

//启动链接，让当前的链接准备工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID= ", c.ConnId)
	//todo 启动从当前链接的读数据的业务
	go c.StartReader()
	for {
		select {
		case <-c.ExitChal:
			return
		}
	}

}

//停止链接，结束当前链接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", c.ConnId)
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Conn Stop() close faild", err)
	}
	c.ExitChal <- true
	close(c.ExitChal)
}

//获取当前链接的绑定sokect conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

//获取远程客户端的TCP状态 IP Port

func (c *Connection) RemoterAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据，将舒服发送给远程的客户端

func (c *Connection) Send(data []byte) error {
	return nil
}
