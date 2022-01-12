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
			//判断邮箱是否被注册
			authGroup.POST("/signup/email/exist",suc.IsEmailExist)
			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}
