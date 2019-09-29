package main

import "fmt"

func main() {
	// 枚举使用 iota

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)


	// iota 遇到 const 就变成 0
	const d = iota
	fmt.Printf("d = %d\n", d)

	// 同一行的 iota 的值一样
	const (
		a1 = iota
		b1, b2, b3 = iota, iota, iota
		c1 = iota
	)

	fmt.Printf("a1 = %d, b1 = %d, b2 = %d, b3 = %d, c1 = %d", a1, b1, b2, b3, c1)
}