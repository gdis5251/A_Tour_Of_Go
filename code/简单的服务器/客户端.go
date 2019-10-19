package main

import (
	"fmt"
	"net"
)

func main() {
	// 连接服务器
	conn, dialErr := net.Dial("tcp", "127.0.0.1:9090")
	if dialErr != nil {
		fmt.Println("dial error = ", dialErr)
		return
	}
	// 给服务器发送消息
	for true {
		fmt.Printf("请输入要发送的文字：")
		var msg string
		_, err := fmt.Scanf("%s", &msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("I say : ", msg)
		conn.Write([]byte(msg[:len(msg)]))
	}

}
