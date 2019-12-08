package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello_json", func(context *gin.Context) {
		fullPath := context.FullPath()

		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "success",
			"data":    fullPath,
		})
	})

	engine.GET("/hello_json_2", func(context *gin.Context) {
		fullPath := context.FullPath()

		resp := Response{
			Code:    1,
			Message: "success",
			Data:    fullPath,
		}

		fmt.Println(resp)

		context.JSON(200, &resp)
	})

	engine.Run()
}

type Response struct {
	Code    int
	Message string
	Data    string
}
