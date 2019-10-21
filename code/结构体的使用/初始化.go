package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
}

func main() {
	// 顺序初始化，每个成员都要初始化
	var s1 Student = Student{1, "Seligman", 'm', 22}
	fmt.Println("s1 = ", s1)

	// 指定成员初始化
	s2 := Student{name: "胡悦", sex: 'f'}
	fmt.Println("s2 = ", s2)

	// 指针的初始化
	var s3 *Student = &Student{1, "Seligman", 'm', 22}
	fmt.Println(s3)

	s4 := &Student{name: "胡悦"}
	fmt.Println(s4)
}
