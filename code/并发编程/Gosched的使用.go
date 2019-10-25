/*
@Time : 2019-10-25 下午 3:08
@Author : Gerald
@File : Gosched的使用
@Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("func go")
		}
	}()

	// Gosched 的作用是让其他协程先执行
	runtime.Gosched()

	for i := 0; i < 2; i++ {
		fmt.Println("main hello")
	}
}
