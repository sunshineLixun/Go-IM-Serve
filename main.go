package main

import (
	"fmt"
	"gin-chat/app/user"
	"gin-chat/routers"
	"gin-chat/utils"
)

func main() {

	utils.InitConfig()
	utils.InitMySQL()

	routers.Include(user.Routers)
		// 初始化路由
		r := routers.Init()
		if err := r.Run(":8081"); err != nil {
			fmt.Println("startup service failed, err:", err)
		}
}