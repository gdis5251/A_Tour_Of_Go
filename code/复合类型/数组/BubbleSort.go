package main

import (
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}
}

func main() {
	// 冒泡排序
	var arr [100]int
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	var s []int = arr[:]

	BubbleSort(s)

	fmt.Println(arr)
}