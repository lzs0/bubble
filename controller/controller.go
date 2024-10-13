package controller

import (
	"bubble/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateTodo(c *gin.Context) {
	// 前端发送的请求
	// 获取前端数据
	var todo models.Todo
	c.BindJSON(&todo)
	// 进行处理
	//...
	// 写入数据库
	if err := todo.CreateTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
		//c.JSON(http.StatusOK, gin.H{
		//	"code": 2000,
		//	"msg": "success",
		//	"data": todo,
		//})
	}
}

func GetTodos(c *gin.Context) {
	sliceTodos, err := models.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, sliceTodos)
	}
}

func GetTodo(c *gin.Context) {

}

func UpdateTodo(c *gin.Context) {
	// 获取url中的id
	id, ok := c.Params.Get("id")
	var todo models.Todo
	if !ok {
		// 没获取到
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	} else {
		// 修改处理
		err := todo.GetATodo(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.BindJSON(&todo)
		if err = todo.UpdateTodo(); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	}
}

func DeleteTodo(c *gin.Context) {
	fmt.Println("delete")
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		fmt.Println("err1")
		return
	} else {
		fmt.Println("ok")
		if err := models.DeleteATodo(id); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{id: "deleted"})
		}
	}
}
