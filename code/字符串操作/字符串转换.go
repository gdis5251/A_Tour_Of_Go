/*
@Time : 2019-10-23 下午 4:31
@Author : Gerald
@File : 字符串转换
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Append 是将其他类型转换为字符串，然后追加在字符串后面
	strSlice := make([]byte, 0, 1024)
	strSlice = strconv.AppendBool(strSlice, true)
	strSlice = strconv.AppendInt(strSlice, 666, 16)
	strSlice = strconv.AppendQuote(strSlice, "I Love HuYue!")
	fmt.Println(string(strSlice))

	// Format 是将其他类型转换为字符串
	var str string
	str = strconv.FormatBool(true)
	fmt.Println(str)
	str = strconv.Itoa(666)
	fmt.Println(str)

	// Parse 是将字符串转换为其他类型
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
