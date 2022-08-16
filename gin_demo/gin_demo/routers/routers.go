package routers

import "github.com/gin-gonic/gin"

func LoadRouter(e *gin.Engine) {
	v1 := e.Group("v1")
	{
		v1.GET("", func(context *gin.Context) {

		})
	}
}
