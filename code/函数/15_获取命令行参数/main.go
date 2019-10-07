package main

import (
	"fmt" 
	"os"
)

func main() {
	list := os.Args
	n := len(list)

	fmt.Println(n)

	for i, data := range list{
		fmt.Printf("list[%d] = %d\n", i, data)
	}
}