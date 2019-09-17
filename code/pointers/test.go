package main

import "fmt"

func main() {
	// var (
	// 	ia int = 1
	// 	fb float64 = 1.0
	// )

	// fmt.Printf("%p, %p\n", &ia, &fb)

	var flag string = "I will hard-working in ByteDance."
	ptr := &flag
	
	fmt.Printf("ptr is %T\n", ptr)
	fmt.Printf("%s\n", *ptr)

	value := *ptr
	fmt.Printf("value is %T\n", value)
	fmt.Printf("%s\n", value)

}