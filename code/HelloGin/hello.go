package main

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"log"
//)
//
//func main() {
//	engine := gin.Default()
//
//	engine.GET("/hello", func(context *gin.Context) {
//		fmt.Println("请求路径: " + context.FullPath())
//		context.Writer.Write([]byte("郭文峰是最棒的！！！\n"))
//	})
//
//	// 运行服务器，若运行出错，打印日志
//	if err := engine.Run(); err != nil {
//		log.Fatal(err.Error())
//	}
//}
