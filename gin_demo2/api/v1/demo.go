package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct {
}

// DemoRegister 负责demo路由注册
func DemoRegister(router *gin.RouterGroup) {
	demo := DemoController{}
	router.GET("/", demo.Index)
	router.GET("/test", demo.Test)
	router.GET("/test2", demo.Test2)
	//router.GET("/dao", demo.Dao)
	//router.GET("/redis", demo.Redis)
}

//以下为路由对应handle

func (controller DemoController) Index(context *gin.Context) {
	//middleware.ResponseSuccess(c, "")
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": "标题",
	})
	return
}
func (controller DemoController) Test(context *gin.Context) {
	// 无意抛出 panic
	var slice = []int{1, 2, 3, 4, 5}
	slice[6] = 6
	return
}
func (controller DemoController) Test2(context *gin.Context) {
	// 主动抛出 panic
	panic("主动panic")
	return
}
