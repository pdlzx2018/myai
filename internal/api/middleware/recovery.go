package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// Recovery 自定义恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 打印堆栈信息
				debug.PrintStack()

				// 记录错误日志
				fmt.Printf("panic recovered: %v\n", err)

				// 返回服务器错误
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "服务器内部错误",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
