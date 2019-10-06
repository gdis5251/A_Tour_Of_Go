package main

import "fmt"

// func test1() int {
// 	var x int
// 	x++
// 	return x * x
// }
// 
// func main() {
// 	fmt.Println(test1())
// 	fmt.Println(test1())
// 	fmt.Println(test1())
// 	fmt.Println(test1())
// 	fmt.Println(test1())
// }

func test() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	// 返回值为一个匿名函数，返回一个函数类型
	// 通过 f 来调用返回的匿名函数
	f := test()
	fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())

}