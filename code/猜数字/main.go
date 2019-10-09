package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 产生一个随机数
	rand.Seed(time.Now().UnixNano()) // 设置一个随机种子
	num := rand.Intn(100)

	for true {
	 	fmt.Printf("请输入一个 100 以内的数字：")
	 	var input int
		fmt.Scanf("%d %c", &input)
		 
	 	if input == num {
	 		fmt.Println("Bingo!")
	 		break
	 	} else if input < num {
	 		fmt.Println("猜小了，再来一次吧！")
	 	} else {
	 		fmt.Println("猜大了，再往小猜猜~")
	 	}
	}
}