package user

import (
	categoryModel "enterprise-api/app/models/category"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	type0 := core.ToInt(c.DefaultQuery("type", "2"))
	list, err := categoryModel.List(type0)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
