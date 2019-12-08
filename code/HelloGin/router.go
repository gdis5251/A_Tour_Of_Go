package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	// 使用 GET 方法来获取用户名
	engine.GET("/show", func(context *gin.Context) {
		// 打印路径
		fmt.Println(context.FullPath())

		// 获取用户名
		name := context.Query("name")

		// 给浏览器打印用户名
		context.Writer.Write([]byte("hello " + name))
	})

	// 使用 POST 方法来进行获取用户名和密码，并给浏览器打印用户名和密码
	engine.POST("/login", func(context *gin.Context) {
		// 打印路径
		fmt.Println(context.FullPath())

		// 获取用户名
		username, existUsername := context.GetPostForm("username")
		if existUsername {
			fmt.Println("username : " + username)
		}

		// 获取密码
		password, existPassword := context.GetPostForm("password")
		if existPassword {
			fmt.Println("password : " + password)
		}

		context.Writer.Write([]byte("username : " + username + "\n" + "password : " + password + "\n"))
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
