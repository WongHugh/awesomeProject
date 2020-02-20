//Time    : 2020-02-20 12:10
//Author  : Hugh
//File    : datapack_test.go
//Descripe:

package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//测试datapack拆包封包的单元测
func TestDataPack(t *testing.T) {
	/*
		模拟服务器
	*/
	listenner, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
	}
	//服务端
	go func() {
		for {
			conn, err := listenner.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read headData error:", err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error:", err)
						return
					}

					if msgHead.GetMsgfLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMsgfLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data error,", err)
							return
						}
						fmt.Println("---->Recv MsgID:", msg.Id, ",dataLen:", msg.DataLen, ",data:", string(msg.Data))
					}
				}

			}(conn)

		}
	}()

	/*
	 模拟客户端
	*/
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	dp := NewDataPack()

	//模拟粘包
	msg1 := &Message{
		Id:      1,
		DataLen: 4,
		Data:    []byte{'h', 'o', 'l', 'l'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err", err)
		return
	}

	msg2 := &Message{
		Id:      2,
		DataLen: 7,
		Data:    []byte{'e', 'g', 'o', 'l', 'a', 'n', 'g'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg2 err", err)
		return
	}
	sendData1 = append(sendData1, sendData2...)
	_, err = conn.Write(sendData1)
	if err != nil {
		fmt.Println(" conn write wrong!,err:", err)
		return
	}

	//客户端阻塞
	select {}

}
