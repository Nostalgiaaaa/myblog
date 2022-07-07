package main

import (
	"myblog/model"
	"myblog/routes"
	"myblog/utils"
)

func main() {

	//viper 读取
	utils.Init()

	//连接数据库
	model.InitDb()

	routes.InitRouter()
}
