package main

import "fmt"

func Swap(lhs *int, rhs *int) {
	*lhs, *rhs = *rhs, *lhs
}

func main() {
	// var igirl int = 20
	// var pgirl *int = & igirl

	// fmt.Println(igirl)
	// fmt.Println(pgirl)


	// // 指针的操作
	// *pgirl = 921
	// fmt.Println(igirl)


	a, b := 10, 20
	Swap(&a, &b)

	fmt.Println(a, b)
}