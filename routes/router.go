package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qiublog/middleware"
	"qiublog/utils"
	"qiublog/utils/sitemap"
	"strings"
)

func InitRouter() {
	gin.SetMode(utils.Config.Server.AppMode)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.Use(gin.Recovery())
	r.Use(middleware.RateMiddleware()) // 胜率限制
	r.Use(middleware.Cors())
	loadStatic(r)
	startAuth(r)
	startUser(r)
	err := r.Run(utils.Config.Server.HttpPort)
	if err != nil {
		panic(fmt.Sprintf("Server startup failure, %v", err))
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
