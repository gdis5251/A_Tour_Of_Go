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

	// fmt.Println(split(17))

	// 循环
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("I will join in ByteDance be a bytedancer.", i)
	// }

	// 使用 for 来替代 C 中的 while
	// count := 10
	// for count >= 0 {
	// 	fmt.Println(count)
	// 	count--;
	// }

	// 无限循环
	// for {
	// 	 不需要写条件，只写一个 for 是无限循环
	// }

	// if-else 语句
	// ByteDancer := true 
	// if ByteDancer {
	// 	fmt.Println("congratulation!")
	// } else {
	// 	fmt.Println("It is alright, TW still pretty good! And congratulation!")
	// }

	// defer 关键字
	// 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	// defer fmt.Println("Go, believe yourself!")
	// fmt.Println("I'd like to join ByteDance!")

	// defer 栈
	fmt.Println("counting:")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done!")
}