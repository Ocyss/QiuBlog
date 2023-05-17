package main

import (
	"qiublog/db"
	"qiublog/model"
	"qiublog/routes"
	"qiublog/utils/log"
	"qiublog/utils/sitemap"
)

func main() {
	log.InitLogger()
	model.InitDb()
	db.InitRedis()
	sitemap.InitSitemap()
	routes.InitRouter()
}
