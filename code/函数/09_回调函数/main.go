package main

import "fmt"

type FuncType func(int, int) int 

// 回调函数：函数参数是函数类型
func Cal(x int, y int, fun FuncType) (result int) {
	result = fun(x, y) 
	return
}

func Add(lhs int, rhs int) (result int) {
	result = lhs + rhs
	return
}

func Min(lhs int, rhs int) (result int) {
	result = lhs - rhs
	return 
}


func main() {
	result := Cal(1, 2, Add)
	fmt.Println(result)

	result = Cal(4, 2, Min)
	fmt.Println(result)
}