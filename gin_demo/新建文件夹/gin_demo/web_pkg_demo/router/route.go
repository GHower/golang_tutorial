package router

import (
	"fmt"
	"gin_demo/web_pkg_demo/controller"
	"gin_demo/web_pkg_demo/docs"
	"gin_demo/web_pkg_demo/lib"
	"gin_demo/web_pkg_demo/middleware"
	"gin_demo/web_pkg_demo/models/dao"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	// swagger配置
	docs.SwaggerInfo.Title = lib.GetStringConf("base.swagger.title")
	docs.SwaggerInfo.Description = lib.GetStringConf("base.swagger.desc")
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = lib.GetStringConf("base.swagger.host")
	docs.SwaggerInfo.BasePath = lib.GetStringConf("base.swagger.base_path")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	//
	router := gin.New()
	// 这里use的主要是一些全局的，外部的中间件
	router.Use(middlewares...)
	//router.LoadHTMLFiles("templates/index.html",)
	router.LoadHTMLGlob("./gin_scaffold/templates/*")
	// 普通路由
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// sagger路径
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 其他简单路由或路由组
	// ---------------------------------------------------------------------
	// 不限流测试
	router.GET("/login", loginTest)
	//router.Use(middleware.LimitingApi()).GET("/login", loginTest)

	// 普通限流，会争抢有限资源，最大资源数有限制，但请求结束后会释放
	v1 := router.Group("v1")
	v1.Use(middleware.LimitingApi())
	{
		v1.GET("/login", loginTest)
	}
	// 通道模拟的令牌桶，类似放票的模式
	v2 := router.Group("v2")
	v2.Use(middleware.LimitingApiTokenBucket())
	{
		v2.GET("/login", loginTest)
	}

	// 通道模拟的漏桶
	v3 := router.Group("v3")
	v3.Use(middleware.LimitingApiLeakyBucket())
	{
		v3.GET("/login", loginTest)
	}

	// demo
	demo := router.Group("demo")
	// 使用中间件
	demo.Use(middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.IPAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.DemoRegister(demo)
	}

	// 需要权限、恢复、日志等中间件的路由接口
	// 非登录接口
	store := sessions.NewCookieStore([]byte("secret"))
	apiNormalGroup := router.Group("/api")
	apiNormalGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.TranslationMiddleware())
	{
		controller.ApiRegister(apiNormalGroup)
	}
	// 登录接口
	apiAuthGroup := router.Group("/api/auth")
	apiAuthGroup.Use(
		sessions.Sessions("mysession", store),
		middleware.RecoveryMiddleware(),
		middleware.RequestLog(),
		middleware.SessionAuthMiddleware(),
		middleware.TranslationMiddleware())
	{
		controller.ApiLoginRegister(apiAuthGroup)
	}
	return router
}

func loginTest(ctx *gin.Context) {
	dbName := ctx.GetHeader("db")
	fmt.Println("db====", dbName)

	var list []dao.User

	// 查询数据，传入
	db, err := lib.GetGormPool(dbName)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}

	//db.Debug().AutoMigrate(&dao.User{})

	db.Find(&list)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "标题",
		"list":  list,
	})
}
