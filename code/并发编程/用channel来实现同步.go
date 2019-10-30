/*
@Time : 2019-10-28 下午 5:30
@Author : Gerald
@File : 用channel来实现同步
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func printString(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
	}
	fmt.Printf("\n")
}

// var ch = make(chan int)

// 使用 sync.WaitGroup 的方式来实现主协程等待其他子协程
var wg sync.WaitGroup

var tongBu = make(chan int)

func person1() {
	printString("Gerald")
	tongBu <- 1
	// ch <- 1

	wg.Done()

}

func person2() {
	<-tongBu
	printString("Seligman")
	// ch <- 2

	wg.Done()
}

func main() {
	wg.Add(2)

	// 目的：使用 channel 来实现 person1 先于 person2 执行
	go person1()

	go person2()
	defer close(tongBu)

	//count := 2
	//
	//// 判断所有协程是否退出
	//for range ch {
	//	count--
	//
	//	if 0 == count {
	//		close(ch)
	//	}
	//}

	wg.Wait()
}
