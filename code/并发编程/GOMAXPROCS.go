/*
@Time : 2019-10-28 下午 5:11
@Author : Gerald
@File : GOMAXPROCS
@Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 返回值是之前默认的有多少核
	// 参数是设置为多少核工作
	n := runtime.GOMAXPROCS(1)
	fmt.Println(n)

	for true {
		go fmt.Print("副")

		fmt.Print("主")
	}
}
