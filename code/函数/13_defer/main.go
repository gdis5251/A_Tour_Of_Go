package main

import "fmt"

func DivZero(x int) {
	result := 100 / x
	fmt.Println(result)
}

func main() {
	// defer 延迟调用，main函数结束前调用、
	// defer fmt.Printf("Hello\n")
	// fmt.Printf("World\n")

	defer fmt.Println("aaaaaaaaa")
	defer fmt.Println("bbbbbbbb")
	defer DivZero(0)
	defer fmt.Println("ccccccccc")
}