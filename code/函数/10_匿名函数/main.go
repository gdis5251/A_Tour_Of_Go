package main

import "fmt"

func main() {
	var a int = 21
	var str string = "Seligman"

	// 匿名函数定义
	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()


	// 定义匿名函数，同时调用
	func() {
		fmt.Println("I am no-name function!")
	}()
}