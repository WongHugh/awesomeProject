//Time    : 2020-02-20 8:58
//Author  : Hugh
//File    : datapack.go
//Descripe:

package znet

import (
	"awesomeProject/src/zinx/utils"
	"awesomeProject/src/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

//初始化实例
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取报的头的长度的方法
func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4字节)+Id uint32(4个字节)
	return 8
}

//封包方法

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	dataBuf := bytes.NewBuffer([]byte{})
	//将dataLen 写入dataBuf
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgfLen()); err != nil {
		return nil, err
	}

	//将MsgId写入dataBuf
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//将data数据写入dataBuf
	if err := binary.Write(dataBuf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuf.Bytes(), nil
}

//拆包方法

func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	dataBuf := bytes.NewReader(binaryData)

	msg := &Message{}

	//读取dataLen
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读MsgId
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	//判断DataLen是否已经超出了允许的最大包程度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil

}
