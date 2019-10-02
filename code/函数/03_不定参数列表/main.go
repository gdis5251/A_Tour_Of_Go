package main

import "fmt"

// 不定参数列表函数
// ...type 不定参数类型
// 传递的实参可以是 0 或多个
func MyFunc(args ...int) {
	fmt.Println("===========================")
	if len(args) == 0 {
		fmt.Println("args is empty!")
		fmt.Println("===========================")
		return
	}
	
	fmt.Println("len(args) = ", len(args))
	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}

	fmt.Println("===========================")
}

func main() {
	MyFunc(1, 2, 3, 4)
	MyFunc(1)
	MyFunc()
}