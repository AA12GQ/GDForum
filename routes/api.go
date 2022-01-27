package routes

import (
	controllers "GDForum/app/http/controllers/api/v1"
	"GDForum/app/http/controllers/api/v1/auth"
	"GDForum/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine){
	v1 := r.Group("/v1")
	v1.Use(middlewares.LimitIP("200-H"))
	{
		authGroup := v1.Group("/auth")
		authGroup.Use(middlewares.LimitPerRoute("1000-H"))
		{
			suc := new(auth.SignupController)
			//判断手机号是否被注册
			authGroup.POST("/signup/phone/exist",middlewares.GuestJWT(),
				suc.IsPhoneExist)
			//判断邮箱是否被注册
			authGroup.POST("/signup/email/exist",middlewares.GuestJWT(),
				suc.IsEmailExist)
			authGroup.POST("/signup/using-phone",middlewares.GuestJWT(),
				suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email",middlewares.GuestJWT(),
				suc.SignupUsingEmail)
			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha",middlewares.LimitPerRoute("50-H"),
				vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone",middlewares.LimitPerRoute("20-H"),
				vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email",middlewares.LimitPerRoute("20-H"),
				vcc.SendUsingEmail)
			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone",middlewares.GuestJWT(),
				lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password",middlewares.GuestJWT(),
				lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token",middlewares.GuestJWT(),
				lgc.RefreshToken)
			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone",middlewares.GuestJWT(),
				pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email",middlewares.GuestJWT(),
				pwc.ResetByEmail)
			uc := new(controllers.UsersController)
			// 获取当前用户
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			userGroup := v1.Group("/users")
			{
				userGroup.GET("",uc.Index)
				userGroup.PUT("",middlewares.AuthJWT(),uc.UpdateProfile)
			}
			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories")
			{
				cgcGroup.GET("",cgc.Index)
				cgcGroup.POST("",middlewares.AuthJWT(),cgc.Store)
				cgcGroup.PUT("/:id",middlewares.AuthJWT(),cgc.Update)
				cgcGroup.DELETE("/:id",middlewares.AuthJWT(),cgc.Delete)
			}
			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics")
			{
				tpcGroup.GET("",middlewares.AuthJWT(),tpc.Index)
				tpcGroup.POST("",middlewares.AuthJWT(),tpc.Store)
				tpcGroup.POST("/:id",middlewares.AuthJWT(),tpc.Update)
				tpcGroup.DELETE("/:id",middlewares.AuthJWT(),tpc.Delete)
				tpcGroup.GET("/:id",tpc.Show)
			}
			lsc := new(controllers.LinksController)
			lscGroup := v1.Group("/links")
			{
				lscGroup.GET("",lsc.Index)
			}
		}
	}
}
