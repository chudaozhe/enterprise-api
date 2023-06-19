package user

import (
	pageModel "enterprise-api/app/models/pages"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func DetailPage(c *gin.Context) {
	var detailPageIn schemas.DetailPageIn
	if err := c.ShouldBindUri(&detailPageIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := pageModel.FindById(detailPageIn.PageId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, detail)
	}
}
