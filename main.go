package main

import (
	"bubble/configs"
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	//创建数据库

	//连接数据库
	dao.SqlInit()
	// defer DB.Close() 自动管理，无需手动管理

	// 迁移表
	dao.DB.AutoMigrate(&models.Todo{})

	//处理路由
	r := routers.SetRouters()

	// run
	if err := r.Run(fmt.Sprintf(":%d", configs.PORT)); err != nil {
		fmt.Println("serve startup failed..")
	}

}
