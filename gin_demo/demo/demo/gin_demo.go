package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("出错了")
		return
	}

}
