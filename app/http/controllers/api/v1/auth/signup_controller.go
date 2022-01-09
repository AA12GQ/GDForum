// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/models/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController)IsPhoneExist(c *gin.Context){
	//请求对象
	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBindJSON(&request);err != nil{
		//解析失败，返回422状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"err":err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsPhoneExist(request.Phone),
	})
}