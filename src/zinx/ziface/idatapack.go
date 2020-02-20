//Time    : 2020-02-20 8:58
//Author  : Hugh
//File    : idatapack.go
//Descripe:

package ziface

/*
	封包
*/

type IDataPack interface {

	//获取报的头的长度的方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (IMessage, error)
}
