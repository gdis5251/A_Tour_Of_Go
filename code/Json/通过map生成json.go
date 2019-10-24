/*
@Time : 2019-10-24 下午 4:35
@Author : Gerald
@File : 通过map生成json
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]interface{})
	m["company"] = "ByteDance"
	m["language"] = []string{"Golang", "C/C++", "Java", "Python"}
	m["salary"] = 22 * 17
	m["Base"] = "BeiJing"

	resJson, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println(string(resJson))
}
