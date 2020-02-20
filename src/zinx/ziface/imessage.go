//Time    : 2020-02-20 8:38
//Author  : Hugh
//File    : imessage.go
//Descripe:

package ziface

/*
	将请求的消息封装到message中，定义抽象的接口
*/

type IMessage interface {
	//获取消息ID
	GetMsgID() uint32
	//获取消息长度
	GetMsgfLen() uint32
	//获取消息内容
	GetData() []byte
	//设置消息ID
	SetMsgId(uint32)
	//设置消息长度
	SetDataLen(uint32)
	//设置消息内容
	SetData([]byte)
}
