package v1

import (
    "GDForum/app/models/link"
    "GDForum/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    response.Data(c, link.AllCached())
}