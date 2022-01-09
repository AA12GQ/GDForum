// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/models/user"
	"GDForum/app/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}
//IsPhoneExist 检测手机号是否被注册
func (sc *SignupController)IsPhoneExist(c *gin.Context){

	request := requests.SignupPhoneExistRequest{}

	if err := c.ShouldBindJSON(&request);err != nil{
		//解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"err":err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	//表单验证
	errs := requests.ValidateSignupPhoneExist(&request,c)
	if len(errs) > 0{
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"errors":errs,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsPhoneExist(request.Phone),
	})
}