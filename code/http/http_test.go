/*
@Time : 2019-11-5 上午 10:49
@Author : Gerald
@File : http_test
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Seligman!"))
	})

	fmt.Println("准备进行阻塞等待")

	http.ListenAndServe("127.0.0.1:8080", nil)

	fmt.Println("并没有阻塞等待！")
}
