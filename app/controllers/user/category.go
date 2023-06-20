package user

import (
	categoryModel "enterprise-api/app/models/category"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var listCategoryIn schemas.ListCategoryIn
	if err := c.ShouldBindQuery(&listCategoryIn); err != nil {
		core.Error(c, 1, core.Translate(err))
		return
	}
	if listCategoryIn.Type == 0 {
		listCategoryIn.Type = 2
	}
	list, err := categoryModel.List(listCategoryIn.Type)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
