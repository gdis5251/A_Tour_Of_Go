package main

import "fmt"

var Add = func(lhs int, rhs int) (result int) {
	result = lhs + rhs
	return
}

func main() {
	fmt.Println(Add(1, 2))
}
