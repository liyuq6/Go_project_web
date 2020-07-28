package controller

import (
	"go_web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//新增待办事项
func AddTodo(c *gin.Context) {
	//前端页面填写待办事项等待点击提交，发送请求过来
	//1.从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todo,
		})
	}
	//3.返回响应，给出添加操作的结果响应
}

//查询所有待办事项
func FindTodo(c *gin.Context) {
	//查询todo这个表的所有的数据
	todoList, err := models.FindAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

//更新id对应的待办事项
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id") //获取参数的id
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id无效"})
		return
	}
	todo, err := models.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)                              //绑定对象todo
	if err = models.UpdateTodo(todo); err != nil { //修改对象
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//删除id对应的待办事项
func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id") //获取对应的id
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "id无效"})
		return
	}
	if err := models.DeleteTodo(id); err != nil { //删除对应id的对象
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
