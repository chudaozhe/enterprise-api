package admin

import (
	pageModel "enterprise-api/app/models/pages"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var listPageIn schemas.ListPageIn
	if err := c.ShouldBind(&listPageIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	count, list, err := pageModel.List(currentCategory.CategoryId, listPageIn.Page, listPageIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": list})
	}
}

func CreatePage(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	var createPageIn schemas.CreatePageIn
	if err := c.ShouldBind(&createPageIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	page := pageModel.Page{
		Title:      createPageIn.Title,
		Content:    createPageIn.Content,
		CategoryId: currentCategory.CategoryId,
		Image:      createPageIn.Image,
	}
	id, err := page.CreatePage()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}

func DetailPage(c *gin.Context) {
	var currentPage schemas.CurrentPage
	if err := c.ShouldBindUri(&currentPage); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := pageModel.FindById(currentPage.PageId)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, detail)
	}
}

func ChangePage(c *gin.Context) {
	var changePageHeaderIn schemas.ChangePageHeaderIn
	if err := c.ShouldBindUri(&changePageHeaderIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changePageIn schemas.ChangePageIn
	if err := c.ShouldBind(&changePageIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	page, err := pageModel.FindById(changePageHeaderIn.PageId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	page.Title = changePageIn.Title
	page.Image = changePageIn.Image
	page.CategoryId = changePageHeaderIn.CategoryId
	page.Content = changePageIn.Content
	err2 := page.UpdatePage()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DeletePage(c *gin.Context) {
	var currentPage schemas.CurrentPage
	if err := c.ShouldBindUri(&currentPage); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := pageModel.DeleteById(currentPage.PageId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"delete": true})
	}
}
