/*
@Time : 2019-10-25 下午 2:40
@Author : Gerald
@File : goroutine的创建
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func newTask() {
	for true {
		fmt.Println("This is new task......")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask()

	for true {
		fmt.Println("This is main task!")
		time.Sleep(time.Second)
	}
}
