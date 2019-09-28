package main

import "fmt"

type X struct {
	a int
}

type Y struct {
	X
	b int
}

type Z struct {
	Y
	c int
}

func main() {
	x := X{a : 1}
	y := Y{
		X : x,
		b : 2,
	}	
    z := Z{
		Y : y,
		c : 3,
	}


	fmt.Println(z.a, z.b, z.c)

}