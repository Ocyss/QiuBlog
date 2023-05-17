package main

import (
	"qiublog/db"
	"qiublog/model"
	"qiublog/routes"
	"qiublog/utils"
	"qiublog/utils/sitemap"
)

func main() {
	utils.InitLog()
	model.InitDb()
	db.InitRedis()
	sitemap.InitSitemap()
	routes.InitRouter()
}
