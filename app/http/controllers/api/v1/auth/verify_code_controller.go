package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/pkg/captcha"
	"GDForum/pkg/logger"
	"GDForum/pkg/response"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIController
}

//ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController)ShowCaptcha(c *gin.Context){
	id,b64s,err := captcha.NewCaptcha().GenerateCaptcha() //生成验证码
	logger.LogIf(err)
	response.JSON(c,gin.H{
		"captcha_id" : id,
		"captcha_image" :b64s,
	})
}
