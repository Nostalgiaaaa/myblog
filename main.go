package main

import (
	"myblog/model"
	"myblog/routes"
	"myblog/utils"
)

func main() {

	//viper 读取
	utils.Init()

	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
