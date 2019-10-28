# Golang：如何优雅的让所有子协程执行完后再执行主协程

### 方法一：channel 实现同步

使用 channel 来完成同步功能。

```go
/*
@Time : 2019-10-28 下午 5:30
@Author : Gerald
@File : 用channel来实现同步
@Software: GoLand
*/
package main

import (
	"fmt"
)


func printString(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
	}
	fmt.Printf("\n")
}

var ch = make(chan int)
var tongBu = make(chan int)

func person1() {
	printString("Gerald")
	tongBu <- 1
	ch <- 1
}

func person2() {
	<- tongBu
	printString("Seligman")
	ch <- 2
}


func main() {
	// 目的：使用 channel 来实现 person1 先于 person2 执行
	go person1()

	go person2()

	count := 2

	// 判断所有协程是否退出
	for range ch {
		count--

		if 0 == count {
			close(ch)
		}
	}

}
```

- count 表示有所少个协程
- ch 用来子协程与主协程之间的同步
- tongBu 用来两个协程之间的同步
- 主协程阻塞等待数据，每当一个子协程执行完后，就会往 ch 里面写一个数据，主协程收到后会使 count--，当 count 减为 0，关闭 ch，主协程将不阻塞在 range ch。

### 方法二：sync.WaitGroup

Go 语言提供一个更简单的方式就是，sync.WaitGroup 来实现等待。

sync.WaitGroup 内部是实现了一个计数器，它有三个方法

- Add() 用来设置一个计数
- Done() 用来在操作结束时调用，使计数减1
- Wait() 用来等待所有的操作结束，即计数变为0。

```go
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

// 使用 sync.WaitGroup 的方式来实现主协程等待其他子协程
var wg sync.WaitGroup

var tongBu = make(chan int)

func person1() {
   printString("Gerald")
   tongBu <- 1
 
   wg.Done()
}

func person2() {
   <- tongBu
   printString("Seligman")

   wg.Done()
}


func main() {
   wg.Add(2)

   // 目的：使用 channel 来实现 person1 先于 person2 执行
   go person1()

   go person2()
   defer close(tongBu)

   wg.Wait()
}
```

