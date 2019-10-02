package main

import "fmt"

func MyFunc() int {
	return 921
}

// 给返回值命名==》 推荐写法
func MyFunc01() (result int) {
	result = 921
	return
}

func main() {
	a := MyFunc01()
	fmt.Println(a)
}