package middlewares

import (
	"GDForum/app/models/user"
	"GDForum/pkg/config"
	"GDForum/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"GDForum/pkg/jwt"
)

func AuthJWT()gin.HandlerFunc{
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims ,err := jwt.NewJWT().ParserToken(c)
		//JWT解析出错
		if err != nil {
			response.Unauthorized(c,fmt.Sprintf("请查看 %v 相关的接口认证文档"),
				config.GetString("app.name"))
			return
		}
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0{
			response.Unauthorized(c,"找不到对应用户，用户可能已经被删除")
			return
		}
		//将用户信息存入gin.context里，后续auth包将从这里拿到用户数据
		c.Set("current_user_id",userModel.GetStringID())
		c.Set("current_user_name",userModel.Name)
		c.Set("current_user",userModel)

		c.Next()
	}
}
