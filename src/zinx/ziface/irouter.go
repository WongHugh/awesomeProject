//Time    : 2020-02-19 14:43
//Author  : Hugh
//File    : irouter.go
//Descripe:

package ziface

/*
	路由抽象接口
	路由的数据都是IRequest
*/
type IRouter interface {
	//在处理conn业务之前的钩子方法hook
	PreHandle(request IRequest)
	//在处理conn业务的主方法
	Handle(request IRequest)
	//在处理conn业务之后的钩子方法hook
	PostHandle(request IRequest)
}
