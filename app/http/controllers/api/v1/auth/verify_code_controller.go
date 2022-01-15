package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/requests"
	"GDForum/pkg/captcha"
	"GDForum/pkg/logger"
	"GDForum/pkg/response"
	"GDForum/pkg/verifycode"
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
// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.VeifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VeifyCodePhone); !ok {
		return
	}

	// 2. 发送 SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败~")
	} else {
		response.Success(c)
	}
}

// SendUsingEmail 发送 Email 验证码
func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.VeifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VeifyCodeEmail); !ok {
		return
	}

	// 2. 发送 SMS
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败~")
	} else {
		response.Success(c)
	}
}
