package routes

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	v1 "qiublog/api/v1"
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
	r.Use(middleware.RateMiddleware())
	if utils.Config.Server.AppMode == "release" {
		r.Use(static.Serve("/", static.LocalFile("web", true)))
		r.NoRoute(func(c *gin.Context) {
			accept := c.Request.Header.Get("Accept")
			flag := strings.Contains(accept, "text/html")
			if flag {
				content, err := os.ReadFile("web/index.html")
				if (err) != nil {
					c.Writer.WriteHeader(404)
					c.Writer.WriteString("Not Found")
					return
				}
				c.Writer.WriteHeader(200)
				c.Writer.Header().Add("Accept", "text/html")
				c.Writer.Write(content)
				c.Writer.Flush()
			}
		})
	}
	r.StaticFile("config", "./config/config.json")

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

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken(true, 100))
	{
		//文章
		auth.POST("article/add", Handler()(v1.ReleaseArticle))  //发布文章
		auth.PUT("article/:id", Handler()(v1.ModifyArticle))    //更新文章
		auth.DELETE("article/:id", Handler()(v1.DeleteArticle)) //删除文章
		//菜单子项
		auth.POST("menuchild/add", Handler()(v1.AddMenuchild)) //添加菜单子项
		auth.PUT("menuchild/set", Handler()(v1.SetMenuchild))  //设置菜单子项
		//分类
		auth.POST("category/add", Handler()(v1.AddCategory))     //添加分类
		auth.PUT("category/list", Handler()(v1.ModifyCategorys)) //批量修改分类
		//上传
		auth.POST("upload/image", Handler()(v1.Upload)) //上传文件
		//消息
		auth.PUT("message/updata", Handler()(v1.UpMessage))      //更新
		auth.DELETE("message/del", Handler()(v1.DelMessage))     //删除
		auth.PUT("message/reply", Handler()(v1.ReplyQuestion))   //回复
		auth.DELETE("message/clear", Handler()(v1.ClearMessage)) //清除缓存
	}

	router := r.Group("api/v1")
	{
		//文章
		router.GET("article/list", Handler()(v1.GetsArticle)) //获取文章列表
		router.GET("article/:id", Handler()(v1.GetArticle))   //获取单个文章
		//用户
		router.POST("user/register", Handler()(v1.Register)) //注册用户
		router.POST("user/login", Handler()(v1.Login))       //登陆用户
		//菜单子项
		router.GET("menuchilds", Handler()(v1.GetMenuchild))     //获取菜单子项
		router.GET("menuchild", Handler()(v1.GetSingleMenuItem)) //获取单个菜单项
		//分类
		router.GET("category", Handler()(v1.GetCategory)) //获取分类
		//标签
		router.GET("tags", Handler()(v1.GetTags)) //获取全部标签
		//消息
		router.POST("message", Handler()(v1.AddMessage))   //进行留言
		router.POST("question", Handler()(v1.AddQuestion)) //进行提问
		router.GET("message", Handler()(v1.GetMessage))    //获取留言
		router.GET("question", Handler()(v1.GetQuestion))  //获取提问
		//统计
		router.POST("statistics/set/mainuv", Handler()(v1.MainSetUV))
		router.GET("statistics", Handler()(v1.GetStatistics))
	}
	err := r.Run(utils.Config.Server.HttpPort)
	if err != nil {
		return
	}
}
