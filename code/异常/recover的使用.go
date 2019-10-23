/*
@Time : 2019-10-23 下午 3:46
@Author : Gerald
@File : recover的使用
@Software: GoLand
*/
package main

import "fmt"

func visitArr(arr []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	return arr[index]
}

func main() {
	var arr [10]int
	sliceArr := []int(arr[:])

	result := visitArr(sliceArr, 20)
	fmt.Println(result)
}
