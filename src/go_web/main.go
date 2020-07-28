package main

import (
	"go_web/dao"
	"go_web/models"
	"go_web/routers"
)

func main() {
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//关闭数据库连接
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()

	r.Run()
}
