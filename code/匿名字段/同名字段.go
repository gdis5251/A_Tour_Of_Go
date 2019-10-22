package main

import "fmt"

type person struct {
	name string
	sex  byte
	age  int
}

type student struct {
	person // 只有类型，没有名字 =》匿名字段，也就是继承
	id     int
	addr   string
	name   string // 跟 Person 拥有相同的字段
}

func main() {
	var s1 student
	s1.name = "Gerald"
	fmt.Printf("%+v\n", s1)
}
