package main

import "fmt"

func main() {
	a := 666
	str := "Seligman"

	// 闭包内部修改变量，就是修改这个变量本身
	func() {
		a = 921
		str = "Go ByteDance!"
		fmt.Printf("内部：a = %d, str = %s\n", a, str)
	}()
	fmt.Printf("外部：a = %d, str = %s\n", a, str)
}