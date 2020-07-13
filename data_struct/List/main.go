package main

import (
	"fmt"
	"github.com/A_Tour_Of_Go/data_struct/List/list"
)

func main() {
	node1 := list.Node{Data:1, Next:nil}
	l := list.List{}
	l.Append(&node1)
	node2 := list.Node{Data:2, Next:nil}
	node3 := list.Node{Data:3, Next:nil}
	l.Append(&node2)
	l.Append(&node3)

	node4 := list.Node{Data:4, Next:nil}
	node5 := list.Node{Data:5, Next:nil}
	l.Insert(1, &node4)
	l.Insert(0, &node5) // 5 1 4 2 3

	l.Remove(0)
	l.Remove(3)
	l.Remove(1) // 1 2

	node := l.IndexOf(1)
	fmt.Println((*node).Data)
}
