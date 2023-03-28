package main

import (
	"fmt"
	"gin-chat/app/user"
	"gin-chat/docs"
	"gin-chat/routers"
	"gin-chat/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()

	routers.Include(user.Routers)
	// 初始化路由
	r := routers.Init()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := r.Run(":8081"); err != nil {
		fmt.Println("startup service failed, err:", err)
	}
}
