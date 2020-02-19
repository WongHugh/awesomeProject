//Time    : 2020-02-19 16:11
//Author  : Hugh
//File    : globalobj.go
//Descripe:

package utils

import (
	"awesomeProject/src/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

/*
	存储一切有关zinx框架的全局参数，供其它模块使用
	一些参数是可以通过zinx.json由用户进行配置
*/
type GlobalObj struct {
	TcpServer ziface.IRequest
	Host      string
	TcpPort   int
	Name      string

	Version        string //zinx的版本号
	MaxConn        int    //当前服务主机允许的最大连接数
	MaxPackageSize uint32 //数据包最大值

}

/*
	定义一个全局的对外globalobj
*/
var GlobalObject *GlobalObj

/*
	从zinx.json加载用于自定义的参数
*/
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}

}

/*
	提供一个init方法，初始化当前的全局对象
*/
func init() {
	GlobalObject = &GlobalObj{
		TcpServer:      nil,
		Host:           "0.0.0.0",
		TcpPort:        10010,
		Name:           "ZinxServerAPP",
		Version:        "V0.4",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//从配置文件重载全局参数
	GlobalObject.Reload()

}
