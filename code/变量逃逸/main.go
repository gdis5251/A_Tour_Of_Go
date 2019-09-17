package main

import "fmt"

func dummy(a int) int {
	var b int = a
	return b
}

func void() {

}

func main() {
	// 声明变量 c 并打印
	var c int

	void()

	fmt.Println(c, dummy(0))
}