package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {

	// 创建记录日志的文件, os.Stdout 同时将日志写入控制台
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/index", func(c *gin.Context) {
		c.String(200, "Hi, gin ")
	})

	router.Run(":8000")
}
