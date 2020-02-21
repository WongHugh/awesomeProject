//Time    : 2020-02-19 11:28
//Author  : Hugh
//File    : connection.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/utils"
	"awesomeProject/src/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

/*
	链接模块
*/
type Connection struct {
	//当前conn 属于哪个Sever
	TcpServer ziface.IServer
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
	//无缓冲通到，用于读写goroutine之间的消息通讯
	msgChan chan []byte
	//消息的管理msgID 和对应的处理业务API关系
	MsgHandler ziface.IMsgHandle

	//链接属性 集合
	property map[string]interface{}
	//保护链接属性的锁
	propertyLock sync.RWMutex
}

//初始化链接模块的方法

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnId:     connID,
		isClosed:   false,
		MsgHandler: msgHandler,
		msgChan:    make(chan []byte),
		ExitChal:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}
	//将conn加入到ConnManager中
	c.TcpServer.GetConnMgr().Add(c)
	return c
}

//链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("Read goroutine ting is running...")
	defer fmt.Println("connID=", c.ConnId, "Reader is exit,remote addr is ", c.RemoterAddr().String())
	defer c.Stop()
	for {

		//创建一个拆包解包对象
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("Read headData err:", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err:", err)
			break
		}
		var data []byte
		if msg.GetMsgfLen() > 0 {
			data = make([]byte, msg.GetMsgfLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data err:", err)
				break
			}

		}
		msg.SetData(data)

		req := Request{msg: msg, conn: c}
		//从路由中，找到注册绑定的Conn对应的router调用
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandler.DoMsgHandle(&req)

		}

	}

}

/*
	写消息数据至客户端
*/
func (c *Connection) StartWriter() {
	fmt.Println("Writer goroutine is running....")
	defer fmt.Println(c.RemoterAddr().String(), "[conn writer exit!]")
	for {
		select {
		case data := <-c.msgChan:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Println("Send data err:", err)
				return
			}
		case <-c.ExitChal:
			return
		}
	}
}

//启动链接，让当前的链接准备工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID= ", c.ConnId)
	//todo 启动从当前链接的读数据的业务
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallOnConnStart(c)
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
	c.TcpServer.CallOnConnStop(c)
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("Conn Stop() close failed", err)
	}
	c.ExitChal <- true

	c.TcpServer.GetConnMgr().Remove(c)

	close(c.ExitChal)
	close(c.msgChan)
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

// 提供一个SendMsg方法，将数据发送到客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	dp := NewDataPack()
	msg := NewMsgPackage(msgId, data)
	binaryMsg, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack error msg id:", msgId, " error:", err)
	}
	c.msgChan <- binaryMsg

	return nil
}

//设置链接属性

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

//删除链接熟悉
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.property, key)
}
