//Time    : 2020-02-20 14:48
//Author  : Hugh
//File    : msgHandle.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/ziface"
	"fmt"
	"strconv"
)

/*
	消息处理模块的实现
*/
type MsgHandle struct {
	Apis map[uint32]ziface.IRouter
}

//初始化MsgHandle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{make(map[uint32]ziface.IRouter)}
}

//调度执行对应的router消息处理
func (mh *MsgHandle) DoMsgHandle(request ziface.IRequest) {
	//从request中找到msgID
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgID=", request.GetMsgId(), "is NOT Found! Need Register!")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

	//根据msgID 调度对应的router业务
}

//为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//是否已经被注册
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeat api,msgID=" + strconv.Itoa(int(msgId)))
	}
	mh.Apis[msgId] = router
	fmt.Println("Add api msgID=", msgId, " success!")
}
