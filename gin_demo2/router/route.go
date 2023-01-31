package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	v1 "golang_tutorial/gin_demo2/api/v1"
	"golang_tutorial/gin_demo2/middleware"
)

// route.go 负责初始化路由he定义路由组等信息，但不负责具体的路由处理
// 具体路由如何处理由controller层实现，只需要把路由或路由组对象传递给controller
// eg.
//    demo := router.Group("demo")
//    controller.DemoRegister(demo)
// 示例中，创建demo路由组，然后传给controller中的DemoRegister方法，由他实现具体路由
// 即route.go 更多的是负责建立联系

// InitRouter 初始化route，参数middlewares是随初始化一起应用的全局中间件
func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	gin.ForceConsoleColor()
	// html模板路径
	router.LoadHTMLGlob("./gin_demo2/templates/*")

	// 全局中间件
	router.Use(middlewares...)
	// 自定义 全局中间件
	router.Use(middleware.Recover)

	// pprof 接口暴露
	pprof.Register(router)

	// 路由关联，并使用中间件
	demoRouter := router.Group("demo")
	//demoRouter.Use()
	{
		v1.DemoRegister(demoRouter)
	}

	return router
}
