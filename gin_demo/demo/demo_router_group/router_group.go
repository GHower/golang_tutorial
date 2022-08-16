package main

import (
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	ctx.Writer.WriteString("this is a login method")
	//fmt.Println("this is a login method")
}

func main() {
	//初始化router
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("/login", login)
	}

	v2 := router.Group("v2")
	{
		v2.GET("/login", login)
	}

	router.Run()
}
