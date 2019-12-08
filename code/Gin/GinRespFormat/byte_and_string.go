package main

import "github.com/gin-gonic/gin"

// 实现返回格式为 byte 和 string 类型
func main() {
	engine := gin.Default()

	engine.GET("/hello_byte", func(context *gin.Context) {
		fullPath := context.FullPath()

		context.Writer.Write([]byte(fullPath))
	})

	engine.GET("/hello_string", func(context *gin.Context) {
		fullPath := context.FullPath()

		context.Writer.WriteString(fullPath)
	})

	engine.Run()
}
