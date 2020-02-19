//Time    : 2020-02-19 14:12
//Author  : Hugh
//File    : irequest.go
//Descripe:

package ziface

/*
	IRequest接口：
	实际上是把客户端请求的链接系你想和请求的数据封装到了一个Request中
*/

type IRequest interface {
	//得到链接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}
