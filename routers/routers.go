package routers

import (
	"net/http"

	"github.com/Grey0520/isnip_api/controller"
	//_ "github.com/Grey0520/isnip_api/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/Grey0520/isnip_api/logger"
	"github.com/Grey0520/isnip_api/middlewares"

	"github.com/gin-gonic/gin"
	//ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-contrib/pprof"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置成发布模式
	}
	r := gin.New() // 每两秒钟添加一个令牌  全局限流
	// r.Use(logger.GinLogger(), logger.GinRecovery(true),middlewares.RateLimitMiddleware(2*time.Second , 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// r := gin.Default()

	r.LoadHTMLFiles("templates/index.html") // 加载html
	r.Static("/static", "./static")         // 加载静态文件
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 注册swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	v1.POST("/login", controller.LoginHandler)
	v1.POST("/signup", controller.SignUpHandler) // 注册业务路由
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		v1.POST("/snippet", controller.CreateSnippetHandler)
		v1.GET("/snippet", controller.SnippetListHandler)
        v1.POST("/folder",controller.CreateFolderHandler)
        v1.GET("/folder", controller.FolderListHandler)
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		v1.POST("/update_snippet",controller.UpdateSnippetHandler)
	}

	pprof.Register(r) // 注册pprof相关路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
