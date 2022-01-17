package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/models/user"
	"GDForum/app/requests"
	"GDForum/pkg/response"
	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	v1.BaseAPIController
}

func (pc *PasswordController)ResetByPhone( c *gin.Context){
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c,&request,requests.ResetByPhone); !ok{
		return
	}
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0{
		response.Abort404(c)
	}else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

func (pc *PasswordController)ResetByEmail(c *gin.Context){
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c,&request,requests.ResetByEmail); !ok{
		return
	}
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0{
		response.Abort404(c)
	}else{
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
