package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.New()
	// LoggerWithFormatter 中间件会将日志写入 gin.DefaultWriter
	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	// 着色
	gin.ForceConsoleColor()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006-01-02 03:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	router.GET("/index", func(c *gin.Context) {
		c.String(200, "Hi, gin ")
	})

	router.Run(":8000")
}
