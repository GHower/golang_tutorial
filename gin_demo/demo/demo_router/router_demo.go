package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/user/:name/:age", func(context *gin.Context) {
		//从路径获取值
		name := context.Param("name")
		age := context.Param("age")
		a := context.Query("a")
		b := context.DefaultQuery("b", "b default")

		message := name + " | " + age

		context.String(http.StatusOK, "hello %s --- %s  --- b:%s", message, a, b)
	})

	// 运行这个router
	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("出错了")
		return
	}
}
