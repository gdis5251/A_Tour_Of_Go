package main

import "fmt"

type MyGirl struct {
	name string
	age  int
}

func (girl MyGirl) MarryTime() (result string) {
	if girl.age >= 20 {
		result = "You can get marry!"
	} else {
		result = "Hold on! You have to wait a few years!"
	}
	return
}

func main() {
	fairy := MyGirl{"胡悦", 20}
	fmt.Println(fairy.MarryTime())
}
