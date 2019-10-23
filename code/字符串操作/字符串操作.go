/*
@Time : 2019-10-23 下午 4:06
@Author : Gerald
@File : 字符串操作
@Software: GoLand
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains 判断原字符串中是否包含目标字符串
	fmt.Println(strings.Contains("hello world", "world"))
	fmt.Println(strings.Contains("hello", "a"))

	// Join 对一个字符串切片，根据某个规则进行拼接
	strSlice := []string{"Gerald", "Love", "HuYue"}
	fmt.Println(strings.Join(strSlice, "❤"))

	// Index 查看目标字符串在原字符串的对应位置，存在返回下标，不存在返回 -1
	fmt.Println(strings.Index("Hello World", "World"))
	fmt.Println(strings.Index("hello", "a"))

	// Repeat 将一个字符串重复几遍，变成一个新的字符串
	str := "I am Gerald"
	str = strings.Repeat(str, 3)
	fmt.Println(str)

	// Replace 将一个字符串中的某个子串替换为另一个子串
	// 如果 n 小于 0，将替换所有子串
	str = strings.Replace(str, "Gerald", "superman", -1)
	fmt.Println(str)

	// Split 将一个字符串根据某个特定子串将其拆分开，存在一个切片中
	strSlice = strings.Split(str, "I")
	fmt.Println(strSlice)

	// Trim 将一个字符串的前后所有 cutset 去掉
	str = "           I Love U !!!            "
	str = strings.Trim(str, " ")
	fmt.Println(str)

	// Fields 将一个字符串根据空格分割开，存在一个切片中返回
	strSlice = strings.Fields(str)
	fmt.Println(strSlice)
}
