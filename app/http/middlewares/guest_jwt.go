package middlewares

import (
	"GDForum/pkg/jwt"
	"GDForum/pkg/response"
	"github.com/gin-gonic/gin"
)

func GuestJWT()gin.HandlerFunc{
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0{
			_,err := jwt.NewJWT().ParserToken(c)
			if err == nil{
				response.Unauthorized(c,"请使用游客身份登录")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
