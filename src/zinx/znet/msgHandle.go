//Time    : 2020-02-20 14:48
//Author  : Hugh
//File    : msgHandle.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/utils"
	"awesomeProject/src/zinx/ziface"
	"fmt"
	"strconv"
)

/*
	消息处理模块的实现
*/
type MsgHandle struct {
	//存放每个MsgId 所对应的处理方法
	Apis map[uint32]ziface.IRouter
	//负责Worker 取任务的消息队列
	TaskQueue []chan ziface.IRequest
	//业务工作Worker池的Worker数量
	WorkerPoolSize uint32
}

//初始化MsgHandle
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
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

//启动一个Worker工作池
func (mh *MsgHandle) StartWorkerPoll() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

//启动一个Worker工作流程
func (mh *MsgHandle) startOneWorker(workId int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker Id=", workId, " is started")

	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandle(request)
		}
	}
}

//将消息交给TaskQueue, 由worker处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//取余进行负载均衡
	workId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("Add ConnID =", request.GetConnection().GetConnID(), " request MsgID=", request.GetMsgId(),
		" to workID=", workId)
	mh.TaskQueue[workId] <- request
}
