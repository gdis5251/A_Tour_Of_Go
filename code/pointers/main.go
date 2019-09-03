package main

import (
	"fmt"
)

type MyStruct struct {
	x int
	y int
}

var (
	m1 = MyStruct{1, 2}
	m2 = MyStruct{x: 7}
	m3 = MyStruct{}
	p = &MyStruct{7, 14}
)

func main() {
	// 指针的基本操作
	// i, j := 7, 21
	// p := &i
	// fmt.Println(*p)

	// *p = 28
	// fmt.Println(*p)

	// p = &j
	// fmt.Println(*p)

	// 结构体
	// fmt.Println(MyStruct{1, 2})

	// 结构体字段
	// m := MyStruct{1, 2}
	// m.x = 7
	// // 结构体指针
	// p := &m
	// p.y = 14
	// fmt.Println(m)


	// 结构体语法
	// fmt.Println(m1, m2, m3, p)

	// 数组
	// var s [2]string
	// s[0] = "I will join in ByteDance"
	// s[1] = "ThoughtWorks is also pretty good"

	// fmt.Println(s[0])
	// fmt.Println(s[1])
	// fmt.Println(s)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	// 数组的切片，半开半闭区间
	var s []int = arr[1 : 4]
	fmt.Println(s)

	// 切片就像是引用，更改切片元素，原本的元素也会被修改
	s[0] = 100
	fmt.Println(s)
	fmt.Println(arr)
}