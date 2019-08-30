package main

import (
	"fmt"
	_ "math"
)

func add(lhs, rhs int) int {
	return lhs + rhs
}

func Swap(lhs, rhs string) (string, string) {
	return rhs, lhs
}

// 命名返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	// 变量声明
	// var ia int = 1
	// var ib int
	// var ic int8 = 2
	// fa, str := 1.1, "GO"

	// fmt.Println(ia, ib, ic, fa, str)

	// Rand
	// rand.Seed(100)
	// fmt.Println("My favorite number is ", rand.Intn(100))

	// test Printf
	// fmt.Printf("I will be rich, I will make %g money. \n", math.Pow(2, 100))

	// fmt.Println(math.Pi)

	// fmt.Println(add(10, 20))

	// strL, strR := Swap("Make", "Money")
	// fmt.Println(strL, strR)

	fmt.Println(split(17))

}