//Time    : 2020-02-19 10:57
//Author  : Hugh
//File    : Client0.go
//Descripe:

package main

import (
	"awesomeProject/src/zinx/znet"
	"fmt"
	"io"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {
	fmt.Println("client0 start....")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("connect start err:", err)
		return
	}

	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("ZinxV0.8 client0 test message")))
		if err != nil {
			fmt.Println("data pack err:", err)
			return
		}
		_, err = conn.Write(binaryMsg)
		if err != nil {
			fmt.Println("conn write err:", err)
			return
		}

		//读取
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head err:", err)
			break
		}
		MsgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msghead err", err)
			break
		}
		if MsgHead.GetMsgfLen() > 0 {
			msg := MsgHead.(*znet.Message)
			msg.Data = make([]byte, msg.DataLen)
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err", err)
				break
			}
			fmt.Println("---->Recv Server Msg:ID=", msg.Id, " msgLen=", msg.DataLen, " data=", string(msg.Data))

		}
		time.Sleep(1 * time.Second)
	}

}
