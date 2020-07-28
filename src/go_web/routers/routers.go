package routers

import (
	"go_web/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//告诉gin框架静态文件去哪里找
	r.Static("/static", "static")
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)
	//路由组op
	opGroup := r.Group("v1")
	{
		//待办事项
		//添加
		opGroup.POST("/todo", controller.AddTodo)
		//查看所有待办事项
		opGroup.GET("/todo", controller.FindTodo)
		//修改某一个待办事项
		opGroup.PUT("/todo/:id", controller.UpdateTodo)
		//删除某一个待办事项
		opGroup.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
