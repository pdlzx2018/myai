package middleware

import "github.com/gin-gonic/gin"

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: 实现JWT认证逻辑
		c.Next()
	}
}
