package routers

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetRouters() (r *gin.Engine) {
	r = gin.Default()

	// 告诉gin去哪里找模板
	r.LoadHTMLGlob("templates/*")
	// 告诉gin去哪里找加载文件
	r.Static("./static", "static")

	r.GET("/", controller.IndexHandler)
	v1Group := r.Group("/v1")
	{
		// 获取全部表单
		v1Group.GET("/todo", controller.GetTodos)
		// 新建一个表单
		v1Group.POST("/todo", controller.CreateTodo)
		// 修改某一个表单
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		// 删除某一个表单
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return
}
