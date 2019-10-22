package main

import (
	"fmt"
)

type Person struct {
	name string
	age int8
	sex byte
}

func (adult Person) Who() {
	fmt.Println("I am Adult!")
}

type Student struct {
	Person
	id         int
	schoolName string
}

func (student Student) Who() {
	fmt.Println("I am student!")
}


func main() {
	var person Student
	person.Who() // 隐藏了 Person 中的 Who 方法，而使用 Student 中的 Who 方法

	// 想使用 Person 中的 Who 方法，必须显示调用
	person.Person.Who()
}

