//Time    : 2020-02-18 16:54
//Author  : Hugh
//File    : iserver.go
//descripe:

package ziface

//定义一个服务器接口
type IServer interface {
	Start()
	Stop()
	Serve()
}
