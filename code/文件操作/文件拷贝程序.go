/*
@Time : 2019-10-25 下午 1:51
@Author : Gerald
@File : 文件拷贝程序
@Software: GoLand
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 检查参数是否正确
	args := os.Args
	if len(args) != 3 {
		fmt.Println("参数出错！应该是 xxx file1 file2.")
		return
	}

	// 检查是否是同名文件
	srcFileName := args[1]
	destFileName := args[2]
	if srcFileName == destFileName {
		fmt.Println("文件同名，请重新运行程序并注意不要使用同名文件！")
		return
	}

	// 打开 file1 文件
	sF, openError := os.Open(srcFileName)
	if openError != nil {
		fmt.Println("打开 ", srcFileName, " 失败, error: ", openError)
		return
	}

	// 创建 file2 文件
	dF, createError := os.Create(destFileName)
	if createError != nil {
		fmt.Println("创建 ", destFileName, " 失败, error: ", createError)
		return
	}

	// 使用完毕后，关闭文件描述符
	defer sF.Close()
	defer dF.Close()

	// 一行一行的读取源文件，一行一行的写入目标文件
	buffer := make([]byte, 4*1024)
	for true {
		n, readError := sF.Read(buffer)
		if readError != nil {
			// 检查是否已经读完
			if readError == io.EOF {
				break
			} else {
				fmt.Println("read file error: ", readError)
			}
		}

		// 将 buffer 里的数据先写入目标文件中
		dF.Write(buffer[:n])
	}
	fmt.Println("Copy Success!")
}
