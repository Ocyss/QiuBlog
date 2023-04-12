package main

import (
	"qiublog/db"
	"qiublog/model"
	"qiublog/routes"
	"qiublog/utils/sitemap"
)

func main() {

	model.InitDb()
	db.InitRedis()
	sitemap.InitSitemap()
	routes.InitRouter()
}
