/*
@Time : 2019-10-28 下午 8:06
@Author : Gerald
@File : 用channel来实现数据传输
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("len: ", len(ch), "cap: ", cap(ch))

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("子协程: ", i, "len: ", len(ch), "cap: ", cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("主协程: ", num)
	}
}
