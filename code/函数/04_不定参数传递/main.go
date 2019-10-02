package main

import "fmt"

func my_func(temp ...int) {
	for i, data := range temp {
		fmt.Printf("temp[%d] = %d\n", i, data)
	}
}

func test(args ...int) {
	// 将全部参数传递给 my_func
	// my_func(args...)

	// 传递部分元素（切片）
	my_func(args[2:]...) // 从下标为2的元素开始，到最后一个元素，传递过去
	my_func(args[:2]...) // 从下标为0的元素开始，到下标为2的元素的前一个元素，传递过去
}

func main() {
	test(1, 2, 3, 4)
}