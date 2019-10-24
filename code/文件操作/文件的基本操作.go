/*
@Time : 2019-10-24 下午 6:10
@Author : Gerald
@File : 文件的基本操作
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	// 创建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("Create error : ", err)
		return
	}

	// 使用完文件之后，关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 9; i++ {
		buf = fmt.Sprintf("I Will Be Rich %d\n", i)

		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("Write error : ", err)
			return
		} else {
			fmt.Println("Success writen ", n, "byte")
		}
	}
}

func ReadFile(path string) {
	// 先打开文件
	f, openError := os.Open(path)
	if openError != nil {
		fmt.Println("Open file error: ", openError)
		return
	}

	// 使用完之后关闭文件
	defer f.Close()

	readBuffer := make([]byte, 1024)

	n, readError := f.Read(readBuffer)
	if readError != nil {
		fmt.Println("Read file error: ", readError)
		return
	}

	fmt.Println(string(readBuffer[:n]))
}

func main() {
	// 定义一个文件路径
	path := "./文件操作/demo.txt"

	// WriteFile(path)
	ReadFile(path)
}
