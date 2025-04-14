package main

import (
	"Blog/model"
	"Blog/routers"
)

func main() {
	// 引用数据库
	model.InitDb()

	routers.InitRouter()

}
