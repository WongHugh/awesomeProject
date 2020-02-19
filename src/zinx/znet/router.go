//Time    : 2020-02-19 14:43
//Author  : Hugh
//File    : router.go
//Descripe:

package znet

import "awesomeProject/src/zinx/ziface"

type BaseRouter struct {
}

//在处理conn业务之前的钩子方法hook
func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}

//在处理conn业务的主方法
func (br *BaseRouter) Handle(request ziface.IRequest) {

}

//在处理conn业务之后的钩子方法hook
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
