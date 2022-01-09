package routes

import (
	"GDForum/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine){
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			//判断手机号是否被注册
			authGroup.POST("/signup/phone/exist",suc.IsPhoneExist)
		}
	}
}
