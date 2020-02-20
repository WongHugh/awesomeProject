//Time    : 2020-02-19 10:57
//Author  : Hugh
//File    : Client0.go
//Descripe:

package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/
func main() {
	fmt.Println("client start....")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:10010")
	if err != nil {
		fmt.Println("coonect start err:", err)
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Zinx V0.2"))
		if err != nil {
			fmt.Println("wright conn err", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back %s\n", buf[:cnt])
		time.Sleep(1 * time.Second)
	}

}
