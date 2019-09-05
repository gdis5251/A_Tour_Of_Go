package main

import (
	"fmt"
)

// type MyStruct struct {
// 	x int
// 	y int
// }
// 
// var (
// 	m1 = MyStruct{1, 2}
// 	m2 = MyStruct{x: 7}
// 	m3 = MyStruct{}
// 	p = &MyStruct{7, 14}
// )

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

	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(arr)
	// // 数组的切片，半开半闭区间
	// var s []int = arr[1 : 4]
	// fmt.Println(s)

	// // 切片就像是引用，更改切片元素，原本的元素也会被修改
	// s[0] = 100
	// fmt.Println(s)
	// fmt.Println(arr)

	// 切片默认下限是0，上限是该切片的长度
	// var s []int = arr[:2]
	// fmt.Println(s)

	// var s2 []int = arr[1 :]
	// fmt.Println(s2)

	// 切片的长度和容量
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// var s1 []int = arr[:0]
	// printSlice(s1)

	// s1 = s1[:4]
	// printSlice(s1)

	// s1 = s1[:2]
	// printSlice(s1)


	// 切片的零值
	// var s []int
	// printSlice(s)

	// if s == nil {
	// 	fmt.Println("Nil!")
	// }

	// for 循环遍历切片时，每次都会返回两个值，第一个为下标，第二个为下标对应的值
	arr := []int {1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d   %v\n", len(s), cap(s), s)
}