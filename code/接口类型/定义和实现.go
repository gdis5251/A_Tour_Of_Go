package main

import "fmt"

type FuncSet interface {
	BuyTicket()
}

type Adult struct {
	name string
	age  int8
}

func (person *Adult) BuyTicket() {
	fmt.Println("成人全价！")
}

type Student struct {
	Adult
	id int
}

func (person *Student) BuyTicket() {
	fmt.Println("学生半价票！")
}

type Soldier string

func (person *Soldier) BuyTicket() {
	fmt.Println("军人免费！")
}

func BuyTicket(person FuncSet) {
	person.BuyTicket()
}

func main() {
	// 先声明一个接口类的变量
	// var f FuncSet

	// 声明不同的自定义类型调用，使接口类调用同名函数，实现多态
	var adult Adult = Adult{"James", 39}
	// f = &adult
	// f.BuyTicket()

	student := Student{Adult{"Seligman", 22}, 1}
	// f = &student
	// f.BuyTicket()

	soldier := Soldier("锅盖头")
	// f = &soldier
	// f.BuyTicket()

	// 多态
	BuyTicket(&adult)
	BuyTicket(&student)
	BuyTicket(&soldier)
}
