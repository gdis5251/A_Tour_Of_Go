package main

import "fmt"

func MyFunc() (a int, b int, c int) {
	a = 1
	b = 2
	c = 3

	return
}

func main() {
	a, b, c := MyFunc()

	fmt.Println(a, b, c)
}