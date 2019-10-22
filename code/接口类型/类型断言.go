package main

import "fmt"

type Love struct {
	deeply int8
}

func main() {
	typeSlice := make([]interface{}, 3)
	typeSlice[0] = 1
	typeSlice[1] = "Hello World!"
	typeSlice[2] = Love{20}

	// 使用 if.go 来判断类型
	for index, data := range typeSlice {
		if value, ok := data.(int); ok {
			fmt.Printf("typeSlice[%d]'s type is int, and the value is %d\n", index, value)
		} else if value, ok := data.(string); ok {
			fmt.Printf("typeSlice[%d]'s type is string, and the value is %s\n", index, value)
		} else if value, ok := data.(Love); ok {
			fmt.Printf("typeSlice[%d]'s type is Love, and the value is %v\n", index, value)
		} else {
			fmt.Println("Unknown type!")
		}
	}

	// 使用 switch.go 来判断
	for index, data := range typeSlice {
		switch value := data.(type) {
		case int:
			fmt.Printf("typeSlice[%d]'s type is int, and the value is %d\n", index, value)
		case string:
			fmt.Printf("typeSlice[%d]'s type is string, and the value is %s\n", index, value)
		case Love:
			fmt.Printf("typeSlice[%d]'s type is Love, and the value is %v\n", index, value)
		default:
			fmt.Println("Unknown type!")
		}
	}
}
