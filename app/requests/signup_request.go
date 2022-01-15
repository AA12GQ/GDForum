// Package requests 处理请求数据和表单验证
package requests

import (
	"GDForum/app/requests/validators"
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


type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Name            string `valid:"name" json:"name"`
	Password        string `valid:"password" json:"password,omitempty"`
	PasswordConfirm string `valid:"password_confirm" json:"password_confirm,omitempty"`
}


func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"name":             []string{"required", "alpha_num", "between:3,20", "not_exists:users,name"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type SignupUsingEmailRequest struct {
	Email				string `json:"email,omitempty" valid:"email"`
	VerifyCode			string `json:"verify_code,omitempty" valid:"verify_code"`
	Name				string `json:"name,omitempty" valid:"name"`
	Password			string `json:"password,omitempty" valid:"password"`
	PasswordConfirm		string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

func SignupUsingEmail(data interface{},c *gin.Context)map[string][]string{
	rules := govalidator.MapData{
		"email" :					[]string{"required","min:4","max:30","email","not_exists:users,email"},
		"name"  :					[]string{"required","alpha_num","between:3,20","not_exists:users,name"},
		"password" :				[]string{"required","min:6"},
		"password_confirm" :		[]string{"required"},
		"verify_code" :				[]string{"required","digits:6"},
	}
	messages := govalidator.MapData{
		"email" :	[]string{
			"required:Email为必填项",
			"min:Email长度需大于4",
			"max:Eamil长度需小于30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email已被占用",
		},
		"name"  :	[]string{
			"required:Name为必填项",
			"alpha_num:用户格式错误，只允许数字和英文",
			"between:用户名长度需在3~20之间",
		},
		"password":	[]string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
		"password_confirm" : []string{
			"required:确认密码框为必填项",
		},
		"verify_code" :		[]string{
			"required:验证码答案必填",
			"digits:验证码长度必须是 6 位的数字",
		},
	}
	errs := validate(data,rules,messages)

	_data := data.(*SignupUsingEmailRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password,_data.PasswordConfirm,errs)
	errs = validators.ValidateVerifyCode(_data.Email,_data.VerifyCode,errs)

	return errs
}