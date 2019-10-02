package main

import "fmt"

func Sum(num int) int {
	if num == 1 {
		return 1
	}

	return num + Sum(num - 1)
}

func main() {
	sum := Sum(100)
	fmt.Println(sum)
}