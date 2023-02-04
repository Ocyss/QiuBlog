package main

import (
	"qiublog/db"
	"qiublog/model"
	"qiublog/routes"
)

func main() {

	model.InitDb()
	db.InitRedis()
	routes.InitRouter()
}
