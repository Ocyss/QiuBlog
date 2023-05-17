package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"qiublog/utils"
	"qiublog/utils/middleware"
	"qiublog/utils/sitemap"
	"strings"
)

func InitRouter() {
	log.Info("init router...")
	if utils.Config.Server.AppMode == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20                 // 8 MiB
	r.Use(middleware.Logger(log.StandardLogger())) // 使用Logger记录日志
	r.Use(gin.Recovery())                          // 恐慌恢复
	r.Use(middleware.RateMiddleware())             // 速率限制
	r.Use(middleware.Cors())                       // 跨域处理
	loadStatic(r)
	startAuth(r)
	startUser(r)
	log.Info("init router run~")
	err := r.Run(fmt.Sprintf("%s:%s", utils.Config.Server.Host, utils.Config.Server.Port))
	if err != nil {
		log.Panic(fmt.Sprintf("Server startup failure, %v", err))
	}
}

func loadStatic(r *gin.Engine) {

	r.StaticFile("about.md", "./config/about.md")
	r.GET("rss/:type", func(c *gin.Context) {
		var data struct {
			Type string `uri:"type"`
		}
		var response, mimetype string
		var err error

		if err := c.ShouldBindUri(&data); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": err})
			return
		}
		switch strings.ToUpper(data.Type) {
		case "RSS":
			mimetype = "application/rss+xml"
			response, err = sitemap.Feed.ToRss()
		case "ATOM":
			mimetype = "application/atom+xml"
			response, err = sitemap.Feed.ToAtom()
		case "JSON":
			mimetype = "application/feed+json"
			response, err = sitemap.Feed.ToJSON()
		default:
			mimetype = "application/xml"
			response, err = sitemap.Feed.ToSitemap()
		}
		if err != nil {
			c.String(http.StatusServiceUnavailable, "Err...")
			return
		}
		c.Data(http.StatusOK, mimetype, []byte(response))
	})

	r.GET("sitemap.xml", func(c *gin.Context) {
		response, err := sitemap.Feed.ToSitemap()
		if err != nil {
			c.String(http.StatusServiceUnavailable, fmt.Sprintf("msg:%v", err))
			return
		}
		c.Data(http.StatusOK, "application/xml", []byte(response))
	})
}
