package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.POST("add_user", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var user User
		context.BindJSON(&user)

		fmt.Println(user.Id)
		fmt.Println(user.Name)
		fmt.Println(user.Sex)

		context.Writer.Write([]byte(user.Name + " add succsee ! \n"))
	})

	engine.Run()
}

type User struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Id   int    `form:"id"`
}
