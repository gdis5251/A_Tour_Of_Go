package main

import "fmt"

type Student1 struct {
	id   int
	name string
	sex  byte
	age  int
}

func main() {
	var s Student1
	s.id = 1
	s.name = "Seligman"
	s.sex = 'm'
	s.age = 22
	fmt.Println("s = ", s)

	// 指针对结构体成员的操作跟普通成员的操作一样
	var p *Student1 = new(Student1)
	p.id = 2
	p.name = "胡悦"
	p.sex = 'f'
	p.age = 20
	fmt.Println("p->Student = ", p)

	// 比较
	fmt.Println("s == *p", s == *p)

	// 赋值
	*p = s
	fmt.Println(*p)
}
