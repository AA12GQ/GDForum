package auth

import (
	v1 "GDForum/app/http/controllers/api/v1"
	"GDForum/app/requests"
	"GDForum/pkg/auth"
	"GDForum/pkg/jwt"
	"GDForum/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}
	user,err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	}else{
		token := jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}

}
