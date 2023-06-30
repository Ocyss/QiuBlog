package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"qiublog/utils"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{utils.Config.SiteInfo.URL, "http://127.0.0.1:16879", "http://localhost:16879"}
	return cors.New(config)
}
