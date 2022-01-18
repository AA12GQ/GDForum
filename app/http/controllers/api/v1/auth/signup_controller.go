// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/models/user"
	"GDForum/app/requests"
	"GDForum/pkg/response"
	"GDForum/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)
// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}
//IsPhoneExist 检测手机号是否被注册
func (sc *SignupController)IsPhoneExist(c *gin.Context){
	//获取请求参数并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c,&request,requests.ValidateSignupPhoneExist); !ok{
		return
	}

	response.JSON(c,gin.H{
		"exist" : user.IsPhoneExist(request.Phone),
	})
}

//IsEmailExist 检查邮箱是否已被注册
func (sc *SignupController)IsEmailExist(c *gin.Context){
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c,&request,requests.ValidateSignupEmailExist);!ok{
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
	response.JSON(c,gin.H{
		"exist" : user.IsEmailExist(request.Email),
	})
}

//SignupUsingPhone 使用手机和验证码进行注册
func(sc *SignupController)SignupUsingPhone(c *gin.Context){
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c,&request,requests.SignupUsingPhone); ! ok{
		return
	}
	userModel := user.User{
		Name:request.Name,
		Phone:request.Phone,
		Password:request.Password,
	}
	userModel.Create()
	if userModel.ID > 0{
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c,gin.H{
			"token" :token,
			"data" : userModel,
		})
	}else{
		response.Abort500(c,"创建用户失败，请稍后再试~")
	}

}
//SignupUsingEmail 使用邮箱和验证码进行注册
func (sc *SignupController)SignupUsingEmail(c *gin.Context){
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c,&request,requests.SignupUsingEmail); !ok{
		return
	}

	userModel := user.User{
		Name:request.Name,
		Email:request.Email,
		Password:request.Password,
	}
	userModel.Create()

	if userModel.ID > 0{
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c,gin.H{
			"token" :token,
			"data" : userModel,
		})
	}else{
		response.Abort500(c,"创建用户失败，请稍后再试~")
	}
}

