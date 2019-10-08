package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置使当前时间为随机种子

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int() % 100) // 是一个很大的随机数, 可以通过取模运算限定范围
		fmt.Println(rand.Intn(100)) // 使随机数限定在 100 以内
	}
}