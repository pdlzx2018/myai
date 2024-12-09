package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{
				"code": 500,
				"msg":  c.Errors.Last().Error(),
			})
		}
	}
}
