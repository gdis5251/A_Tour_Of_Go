package main

import (
	"fmt"
	"net"
)

// 简单的回显服务器
func main() {
	// 1. 监听
	listener, listenErr := net.Listen("tcp", ":9090")
	if listenErr != nil {
		fmt.Println("listen error = ", listenErr)
		return
	}
	defer listener.Close()

	// 2. 阻塞等待
	for true {
		conn, acceptErr := listener.Accept()
		if acceptErr != nil {
			fmt.Println("accpet error = ", acceptErr)
			continue
		}
		defer conn.Close()

		// 3. 接收客户发来的数据
		for true {
			buffer := make([]byte, 1024)
			n, readErr := conn.Read(buffer)
			if readErr != nil {
				fmt.Println("read error = ", readErr)
			}
			if n == 0 {
				fmt.Println("client out")
				break
			}
			fmt.Println("client said : ", string(buffer[:n]))
		}

		// 测试 goland 将代码提交到 github 上...
		fmt.Println("测试 goland 将代码提交到 github 上...")
		fmt.Println("try again!")
	}
}
