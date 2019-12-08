package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// 绑定结构体和请求中的 body 中的数据
		var register RegisterInfo
		context.ShouldBind(&register)

		fmt.Println(register.Username)
		fmt.Println(register.Password)
		fmt.Println(register.PhoneNumber)

		context.Writer.Write([]byte(register.Username + "login up success!\n"))
	})

	engine.Run()
}

type RegisterInfo struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	PhoneNumber string `form:"phone_number"`
}
