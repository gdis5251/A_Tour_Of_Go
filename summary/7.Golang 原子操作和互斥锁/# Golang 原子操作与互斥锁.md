# Golang 原子操作与互斥锁

**先来看一个代码**

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter  int32
	wg sync.WaitGroup
)

func main() {
	wg.Add(2) // 设置需要等待的 goroutine 的数量为 2

	go addCounter("Gerald")
	go addCounter("Seligman")

	wg.Wait()
	fmt.Println("Final Counter is: %+v", counter)
}

func addCounter(whoAmI string) {
	defer wg.Done()

	for count := 0; count < 2;  count++ {
		value := counter

		// Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
		runtime.Gosched()

		value++

		counter = value
	}
}
```

> Final Counter is:  2

首先这个程序是起了两个 goroutine 。每个 goroutine 都对 counter 进行了两次累加操作，所以理论上 counter 最后应该是 4 而不是 2 。

原因是因为，Gerald 协程获取到 counter 的值后，又让 Seligman 协程去获取 counter 的值。注意此时 counter 的值还没改变，所以他们两个协程都拿到 0 这个值。然后两个协程又将 value++ 后赋值给 counter，等于说做了两次 counter = 1 的操作。第二次循环也是一样的原理。所以最后的值是 2 。

这种情况可以用两中方式避免

- Atomic 原子操作
- Mutex 互斥锁

## Atomic

```go
for count := 0; count < 2;  count++ {
		value := counter

		// Gosched使当前go程放弃处理器，以让其它go程运行。它不会挂起当前go程，因此当前go程未来会恢复执行。
		runtime.Gosched()

		value++

		counter = value
	}
```

这段代码其实有点奇怪，为什么给 counter + 1 要写得这么多步。这其实是在模拟 CPU 内部再给你个值+1的时候需要做的操作。首先寄存器会把 counter 的值拿出来保存在寄存器里，然后在寄存器里进行+1操作，最后再把寄存器里的值放到CPU里（原理是这样，若位置错误，请及时提醒，见谅）。

而 Atomic 就是相当于在 CPU 里加锁，让这三步操作在执行的时候，当前协程不可以被调度切换，等这三步完成之后才可以被调度切换。

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter  int32
	wg sync.WaitGroup
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

		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}
```

> Final Counter is:  4

这次 counter 就变成 4 了，符合预期。

## Mutex

这个应该很熟悉了，互斥锁，就是在对元素操作之前进行加锁，操作完了解锁。就可以保证数据一致了。

```go
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
```

> Final Counter is:  4

叮~:bell:

