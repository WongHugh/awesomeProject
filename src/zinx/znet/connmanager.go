//Time    : 2020-02-21 10:28
//Author  : Hugh
//File    : connmanager.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

/*
	链接管理模块
*/
type ConnManager struct {
	//管理的连接集合
	connections map[uint32]ziface.IConnection
	//保护链接的读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加链接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	connMgr.connections[conn.GetConnID()] = conn
	fmt.Println("connID=", conn.GetConnID(), " add to ConnManager successfully :conn num= ", connMgr.Len())

}

//删除链接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connID=", conn.GetConnID(), " remove from ConnManager successfully :conn num= ", connMgr.Len())

}

//根据connID获取链接
func (connMgr *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()
	if conn, ok := connMgr.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not FOUND!")
	}
}

//得到当前的链接总数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

//清楚并终止所有的连接
func (connMgr *ConnManager) ClearConn() {
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	for connID, conn := range connMgr.connections {
		conn.Stop()
		delete(connMgr.connections, connID)
	}
	fmt.Println("Clear all connections success1 conn num = ", connMgr.Len())

}
