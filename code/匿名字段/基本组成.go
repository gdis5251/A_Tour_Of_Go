package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person // 只有类型，没有名字 =》匿名字段，也就是继承
	id     int
	addr   string
}

func main() {
	// 初始化
	// 顺序初始化
	var s1 Student = Student{Person{"Gerald", 'm', 22}, 1, "china"}
	fmt.Printf("s1 = %+v\n", s1)
	// 指定初始化
	s2 := Student{Person: Person{name: "haha"}, addr: "BeiJing"}
	fmt.Printf("s2 = %+v\n", s2)

	// 成员使用
	s1.name = "apple"
	s1.sex = 'f'
	s1.age = 18
	s1.id = 21
	s1.addr = "TaiWan"

	// 整体赋值
	s1.Person = Person{"yoyo", 'f', 17}

	fmt.Println(s1)
}
