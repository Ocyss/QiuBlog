package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/db"
	"time"
)

func RateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 如果ip请求连接数在两秒内超过5次，返回429并抛出error
		if !db.Allow(c.ClientIP(), 20, 2*time.Second) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			return
		}
		c.Next()
	}
}
