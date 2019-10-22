package main

import "fmt"

func fib() func(n int) int { // 返回一个闭包
	return func(n int) int {
		if n == 0 {
			return 0
		} else if n <= 2 {
			return 1
		} else {
			var pre int = 1
			var mid int = 1
			var last int = 1
			for i := 3; i <= n; i++ {
				last = pre + mid
				pre = mid
				mid = last
			}
			return last
		}
	}
}

func main() {
	f := fib()

	for i := 1; i < 10; i++ {
		fmt.Println(f(i))
	}
}
