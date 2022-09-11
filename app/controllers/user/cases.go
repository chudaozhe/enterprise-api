package user

import (
	casesModel "enterprise-api/app/models/cases"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCase(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	keyword := c.DefaultQuery("keyword", "")
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "10")
	count, list, err := casesModel.List(core.ToInt(categoryId), keyword, false, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": list})
	}
}

func DetailCase(c *gin.Context) {
	id, _ := c.Params.Get("case_id")
	if core.ToInt(id) > 0 {
		detail, err := casesModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
