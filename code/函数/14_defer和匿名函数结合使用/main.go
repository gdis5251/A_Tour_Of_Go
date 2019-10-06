package main

import "fmt"

func main() {
	a := 10
	b := 20

	// defer func() {
	// 	fmt.Printf("a = %d, b = %d\n", a, b)
	// }()

	defer func(a int, b int) {
		fmt.Printf("内部：a = %d, b = %d\n", a, b)
	}(a, b)

	a = 111
	b = 222
	fmt.Printf("外部：a = %d, b = %d\n", a, b)
}