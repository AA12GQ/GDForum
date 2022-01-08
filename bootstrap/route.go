//处理程序初始化逻辑
package bootstrap

import (
	"GDForum/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//路由初始化
func SetupRoute(router *gin.Engine){
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	//  注册 API 路由
	routes.RegisterAPIRoutes(router)

	//  配置 404 路由
	setup404Handler(router)
}
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}
func setup404Handler(router *gin.Engine){
	router.NoRoute(func(c *gin.Context) {
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
}