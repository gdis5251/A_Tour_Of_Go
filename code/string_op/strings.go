package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare 比较字符串大小
	// 相等返回 0 ，前者小返回 -1 ，前者大返回 1
	a := "hello guowenfeng"
	b := "hello bytedance"
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.Compare(a, a))
	fmt.Println(strings.Compare(b, a))

	// EqualFold 忽略大小写，比较两个字符串是否相等
	// 相等返回 true ，不相等返回 false
	fmt.Println(strings.EqualFold("guowenfeng", "GuoWenfeng"))
	fmt.Println(strings.EqualFold("love", "xiaohu"))

}
