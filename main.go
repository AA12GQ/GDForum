package main

import (
	"GDForum/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
