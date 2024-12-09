package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/pdlzx2018/myai/internal/api/handler"
	"github.com/pdlzx2018/myai/internal/api/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New() // 不使用默认中间件

	// 创建限流器：每秒5个请求，最多突发10个请求
	limiter := middleware.NewIPRateLimiter(rate.Limit(5), 10)

	// 全局中间件
	r.Use(middleware.Recovery())         // 自定义恢复中间件
	r.Use(middleware.RateLimit(limiter)) // 限流中间件
	r.Use(middleware.Cors())             // CORS中间件
	r.Use(middleware.ErrorHandler())     // 错误处理中间件

	// API 路由组
	api := r.Group("/api")
	{
		// 认证相关
		auth := api.Group("/auth")
		{
			auth.POST("/login", handler.Login)
			auth.POST("/register", handler.Register)
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.JWTAuth())
		{
			// 聊天相关
			chat := authenticated.Group("/chat")
			{
				chat.POST("/send", handler.SendMessage)
				chat.GET("/history", handler.GetChatHistory)
			}

			// 用户相关
			user := authenticated.Group("/user")
			{
				user.GET("/info", handler.GetUserInfo)
				user.PUT("/info", handler.UpdateUserInfo)
			}
		}
	}

	return r
}
