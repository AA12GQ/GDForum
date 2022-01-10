package main

import (
	"GDForum/bootstrap"
	"flag"
	"fmt"
	btsConig "GDForum/config"
	"GDForum/pkg/config"
	"github.com/gin-gonic/gin"
)
func init(){
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}
func main() {
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	//初始化Logger
	bootstrap.SetupLogger()
	router := gin.New()

	//初始化DB
	bootstrap.SetupDB()
	//初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
