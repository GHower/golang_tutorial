package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//初始化router
	router := gin.Default()
	//定义POST方法路径
	router.POST("/form_post", func(context *gin.Context) {
		//获取传递的值
		name := context.PostForm("name")
		age := context.DefaultPostForm("age", "10")
		//以JSON格式返回
		context.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code": http.StatusOK, "status": "ok"}, "name": name, "age": age})
	})

	router.Run()
}
