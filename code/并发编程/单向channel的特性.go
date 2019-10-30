/*
@Time : 2019-10-30 下午 4:19
@Author : Gerald
@File : 单向channel的特性
@Software: GoLand
*/
package main

func main() {
	ch := make(chan int)

	// 双向 channel 可以隐式转化成单向
	var writeChan chan<- int = ch
	writeChan <- 1

	var readChan <-chan int = ch
	<-readChan

	// 单向无法转化成双向
}
