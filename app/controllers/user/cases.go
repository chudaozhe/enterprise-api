package user

import (
	casesModel "enterprise-api/app/models/cases"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCase(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		_ = c.Error(err)
		return
	}
	var listCaseIn schemas.ListCaseIn
	if err := c.ShouldBindQuery(&listCaseIn); err != nil {
		_ = c.Error(err)
		return
	}
	count, list, err := casesModel.List(currentCategory.CategoryId, listCaseIn.Keyword, false, listCaseIn.Page, listCaseIn.Max)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": list})
	}
}

func DetailCase(c *gin.Context) {
	var detailCaseIn schemas.DetailCaseIn
	if err := c.ShouldBindUri(&detailCaseIn); err != nil {
		_ = c.Error(err)
		return
	}
	detail, err := casesModel.FindById(detailCaseIn.CaseId)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, detail)
	}
}
