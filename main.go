package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger(),gin.Recovery())

	// 注册一个路由
	r.GET("/", func(c *gin.Context) {

		// 以 JSON 格式响应
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})
	//处理404请求
	r.NoRoute(func(c *gin.Context) {
	//获取标头信息Accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString,"text/html"){
			c.String(http.StatusNotFound,"页面返回404")
		}else {
			//默认返回JSON
			c.JSON(http.StatusNotFound,gin.H{
				"error_code" : 404,
				"error_message" : "页面未定义，请确认url和请求方法是否正确",
			})
		}
	})
	// 运行服务，默认为 8080，我们指定端口为 8000
	r.Run(":8000")
}
