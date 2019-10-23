/*
@Time : 2019-10-23 下午 3:27
@Author : Gerald
@File : error接口的应用
@Software: GoLand
*/
package main

import (
	"errors"
	"fmt"
)

func myDiv(numerator int, denominator int) (result float64, err error) {
	err = nil

	// 判断是否发生除 0 操作
	if denominator == 0 {
		err = errors.New("发生除 0 操作！")
	} else {
		result = float64(numerator / denominator)
	}

	return
}

func main() {
	result, err := myDiv(10, 0)
	if err != nil {
		fmt.Println("发生异常：", err)
	} else {
		fmt.Println(result)
	}
}
