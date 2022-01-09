// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func ValidateSignupPhoneExist(data interface{},c *gin.Context)map[string][]string{
	//自定义验证规则
	rules := govalidator.MapData{
		"phone" : []string{"required","digits:11"},
	}
	//自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone" :[]string{
			"required:手机号为必填项，参数名称phone",
			"digits:手机号长度必须为11位的数字",
		},
	}
	//开始验证
	return validate(data,rules,messages)
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func ValidateSignupEmailExist(data interface{},c *gin.Context)map[string][]string{

	rules := govalidator.MapData{
		"email": []string{"required","min:4","max:30","email"},
	}

	messages := govalidator.MapData{
		"email" : []string{
			"required:Email为必填项",
			"min:Email长度需大于4",
			"max:Email长度需小于30",
			"email:Email格式不正确，请提供有效的邮箱地址",
		},
	}
	return validate(data,rules,messages)
}