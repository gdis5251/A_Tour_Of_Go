package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// ShouldBindQuery 是针对有 querystring 设计的
	engine.GET("/login_in", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// 绑定结构体和 querystring 中的数据
		var user UserInfo
		context.ShouldBindQuery(&user)

		fmt.Println("username: " + user.Username)
		fmt.Println("password: " + user.Password)

		context.Writer.Write([]byte("username: " + user.Username + "\nlogin in success!\n"))
	})

	engine.Run()
}

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
