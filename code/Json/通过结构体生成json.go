/*
@Time : 2019-10-23 下午 7:20
@Author : Gerald
@File : 通过结构体生成json
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Network struct {
	Port int
	Ip   []string
}

func main() {
	n := Network{22, []string{"192.1.1.1", "123.123.123.123"}}

	// buf, err := json.Marshal(n)
	buf, err := json.MarshalIndent(n, "", " ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))

}
