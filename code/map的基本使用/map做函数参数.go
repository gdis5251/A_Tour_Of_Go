package main

import "fmt"

func Delete(m map[int]string, index int) {
	delete(m, index)
}

func main() {
	m := map[int]string{1: "郭大宝", 2: "胡小宝", 3: "啦啦啦"}

	// 删除 key 值为 x 的元素
	Delete(m, 3)
	// map 做函数参数是以引用的方式传递的
	fmt.Println("m = ", m)
}
