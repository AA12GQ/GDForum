package auth

import (
	"GDForum/app/models/user"
	"GDForum/pkg/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

func Attempt(email string,password string)(user.User,error){
		userModel := user.GetByMulti(email)
		if userModel.ID == 0{
			return user.User{},errors.New("用户不存在")
		}
		if !userModel.ComparePassword(password){
			return user.User{},errors.New("密码错误")
		}
		return userModel,nil
}

func LoginByPhone(phone string)(user.User,error){
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0{
		return user.User{},errors.New("手机未注册")
	}
	return userModel,nil
}
// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context)user.User{
	userModel,ok := c.MustGet("current_user").(user.User)
	if !ok{
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context)string{
	return c.GetString("current_user_id")
}
