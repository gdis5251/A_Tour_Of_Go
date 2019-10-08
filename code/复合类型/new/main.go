package main

import "fmt"

func main() {
	var p *int
	p = new(int)

	*p = 666
	fmt.Println(*p)

	// Go 语言有 gc，所以不用手动释放空间
}