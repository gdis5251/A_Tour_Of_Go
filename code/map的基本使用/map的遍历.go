package main

import "fmt"

func main() {
	m := map[int]string{1 : "郭大宝", 2 : "胡小宝", 3 : "啦啦啦"}

	for i := 0; i < len(m) + 1; i++ {
		value, ok := m[i]
		if ok {
			fmt.Printf("m[%d] = %s\n", i, value)
		} else {
			fmt.Println("Key 不存在")
		}
	}

	// map 元素的删除
	delete(m, 3)
	fmt.Println("m = ", m)
}
