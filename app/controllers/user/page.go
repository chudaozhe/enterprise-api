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
		core.Error(c, 1, core.Translate(err))
		return
	}
	detail, err := pageModel.FindById(detailPageIn.PageId)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, detail)
	}
}
