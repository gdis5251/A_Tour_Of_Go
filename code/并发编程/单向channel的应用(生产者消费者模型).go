/*
@Time : 2019-10-30 下午 4:28
@Author : Gerald
@File : 单向channel的应用(生产者消费者模型)
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

// 添加一个计数，使子协程全部执行完后，再执行主协程
var wg sync.WaitGroup

func producer(writeCh chan<- int) {
	for i := 1; i < 10; i++ {
		writeCh <- i * i
	}

	// 这里要关闭 writeCh 不然会造成 消费者 一直阻塞读的死锁
	close(writeCh)
	wg.Done()
}

func consumer(readCh <-chan int) {
	for i := range readCh {
		fmt.Println("num is ", i)
	}

	wg.Done()
}

func main() {
	ch := make(chan int)

	wg.Add(2)

	// 生产者消费者模型，一个往 channel 里写，一个从 channel 里读
	go producer(ch)
	go consumer(ch)

	wg.Wait()
}
