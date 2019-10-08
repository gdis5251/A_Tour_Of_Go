package main

import "fmt"

func main() {
	// s1 := []int{}
	// fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	// s2 := make([]int, 5)
	// fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	// s2 = append(s2, 10)
	// fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	// s2 = append(s2, 10)
	// s2 = append(s2, 10)
	// s2 = append(s2, 10)
	// s2 = append(s2, 10)
	// s2 = append(s2, 10)
	// fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[:]
	s[8] = 666
	fmt.Println(arr)
	fmt.Println(s)
	fmt.Printf("&arr[8] = %p, &s[8] = %p\n", &arr[8], &s[8])


	// a := 10
	// fmt.Println(&a)

	s = append(s, 10)
	fmt.Println(s)

	fmt.Printf("&arr[8] = %p, &s[8] = %p, &s[9] = %p\n", &arr[8], &s[8], &s[9])

	s[8] = 921
	fmt.Println(arr)
	fmt.Println(s)

}