package main

import (
	"fmt"
	"sync"
)

var (
	counter  int32
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(2) // 设置需要等待的 goroutine 的数量为 2

	go addCounter("Gerald")
	go addCounter("Seligman")

	wg.Wait()
	fmt.Println("Final Counter is: ", counter)
}

func addCounter(whoAmI string) {
	defer wg.Done()

	for count := 0; count < 2;  count++ {
		//value := counter
		//
		//// Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
		//runtime.Gosched()
		//
		//value++
		//
		//counter = value

		//atomic.AddInt32(&counter, 1)
		//runtime.Gosched()

		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}