package router

import (
	"gin_demo/gin_scaffold/controller"
	"gin_demo/gin_scaffold/docs"
	"gin_demo/gin_scaffold/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// InitRouter 初始化得到一个gin引擎
// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
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

	// 限流测试
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
	// sagger路径
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
	// 其他简单路由组
	// ...

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

	//var list []dao.User

	// 查询数据，传入
	//db, err := lib.GetGormPool("default")
	//if err != nil {
	//	middleware.ResponseError(ctx, 2002, err)
	//	return
	//}

	//db.Debug().AutoMigrate(&dao.User{})

	//db.Find(&list)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "标题",
		//"list":  list,
	})
}
