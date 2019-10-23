/*
@Time : 2019-10-23 下午 3:20
@Author : Gerald
@File : error接口的使用
@Software: GoLand
*/
package main

import "fmt"
import "errors"

func main() {
	err1 := fmt.Errorf("%s", "this is a normal error1")
	fmt.Println(err1)

	err2 := errors.New("this is a normal error2")
	fmt.Println(err2)
}
