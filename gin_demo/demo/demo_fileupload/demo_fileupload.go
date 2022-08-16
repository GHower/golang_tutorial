package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//初始化router
	router := gin.Default()
	//定义POST方法路径
	router.POST("/upload", func(context *gin.Context) {
		//获取传递的值
		name := context.PostForm("name") //根据name属性获取别名
		file, err := context.FormFile("upload")
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": gin.H{"status_code": http.StatusBadRequest, "msg": "a bad request"}})
			return
		}
		filename := file.Filename
		if err := context.SaveUploadedFile(file, filename); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": gin.H{"status_code": http.StatusBadRequest, "msg": "upload file err:" + err.Error()}})
			return
		}
		//以JSON格式返回
		context.JSON(http.StatusOK, gin.H{"status": gin.H{"status_code": http.StatusOK, "status": "ok"}, "name": name, "filename": filename})
	})

	router.Run()
}
