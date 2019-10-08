package main

import "fmt"

func main() {
	// var id [100]int

	// for i := 0; i < len(id); i++ {
	// 	id[i] = i + 1
	// 	fmt.Printf("id[%d] = %d\n", i, id[i])
	// }


	// 数组的初始化
	// 全部初始化
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// 部分初始化
	arr3 := [5]int{1, 2, 3}
	arr4 := [5]int{2:9, 4:21}
	fmt.Println(arr3)
	fmt.Println(arr4)


}