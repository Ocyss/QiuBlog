package main

import (
	"qiublog/model"
	"qiublog/routes"
)

func main() {

	model.InitDb()
	routes.InitRouter()
}
