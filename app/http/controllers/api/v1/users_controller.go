package v1

import (
    "GDForum/app/models/user"
    "GDForum/pkg/auth"
    "GDForum/pkg/response"

    "github.com/gin-gonic/gin"
)

type UsersController struct {
    BaseAPIController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
    userModel := auth.CurrentUser(c)
    response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context){
    data := user.All()
    response.Data(c,data)
}