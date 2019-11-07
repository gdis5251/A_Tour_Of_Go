/*
@Time : 2019-11-7 上午 9:08
@Author : Gerald
@File : CURD.go
@Software: GoLand
*/
package main

import (
	"fmt"
	"github.com/astaxie/goredis"
)

func main() {
	var client goredis.Client

	// 设置要连接服务器的 ip
	client.Addr = "127.0.0.1:6379"

	// 插入数据
	setErr := client.Set("Seligman", []byte("Seligman is a good ByteDancer!"))
	if setErr != nil {
		fmt.Println("set data failed, error : ", setErr)
		return
	}

	// 读取数据
	res, getErr := client.Get("Seligman")
	if getErr != nil {
		fmt.Println("get data failed, error : ", getErr)
		return
	}

	fmt.Println(string(res))
}
