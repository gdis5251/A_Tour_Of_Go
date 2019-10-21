package main

import "fmt"

func main() {
	// map 的定义
	var m1 map[int]string // map[key]value
	fmt.Println("m1 = ", m1)
	// map 只有 len 没有 cap
	fmt.Println("m1 len = ", len(m1))

	// 通过 make 创建 map
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("m2 len = ", len(m2))

	m3 := make(map[int]string, 2)
	m3[0] = "郭老大"
	m3[1] = "胡老二"
	m3[2] = "恩恩爱爱一辈子"
	m3[10] = "还有下面的好几倍器"
	fmt.Println("m3 = ", m3)
	fmt.Println("m3 len = ", len(m3))

	// map 的初始化
	m4 := map[int]string{4:"haha", 2:"李光洙", 1:"刘在石"}
	fmt.Println("m4 = ", m4)
}
