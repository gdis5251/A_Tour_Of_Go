package main

import "fmt"

func Add(lhs int, rhs int) (result int) {
	result = lhs + rhs
	return
}

type myfunc func(int, int) int

func main() {
	var a myfunc
	a = Add
	result := a(1, 2)
	fmt.Println(result)
}