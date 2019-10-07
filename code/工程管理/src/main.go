package main

import (
	"fmt"
	"./calc"
)

func main() {
	result := calc.Add(10, 20)
	fmt.Println(result)
}