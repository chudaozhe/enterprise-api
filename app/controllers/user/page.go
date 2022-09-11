package user

import (
	pageModel "enterprise-api/app/models/pages"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func DetailPage(c *gin.Context) {
	id, _ := c.Params.Get("page_id")
	if core.ToInt(id) > 0 {
		detail, err := pageModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
