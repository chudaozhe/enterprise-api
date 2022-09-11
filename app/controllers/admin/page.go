package admin

import (
	pageModel "enterprise-api/app/models/pages"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListPage(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "100")
	count, list, err := pageModel.List(core.ToInt(categoryId), core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": list})
	}
}

func CreatePage(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	if len(title) > 0 && len(content) > 0 {
		page := pageModel.Page{
			Title:      title,
			Content:    content,
			CategoryId: core.ToInt(categoryId),
			Image:      c.DefaultPostForm("image", ""),
		}
		id, err := page.CreatePage()
		if err != nil {
			core.Error(c, 1, err.Error())
			return
		}
		core.Success(c, 0, gin.H{
			"id": id,
		})
	} else {
		core.Success(c, 400, "参数错误")
	}
}
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

func ChangePage(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	id, _ := c.Params.Get("page_id")
	if core.ToInt(id) > 0 {
		updateData := pageModel.Page{
			Title:      c.DefaultPostForm("title", "-isnil-"),
			Content:    c.DefaultPostForm("content", "-isnil-"),
			CategoryId: core.ToInt(categoryId),
			Image:      c.DefaultPostForm("image", "-isnil-"),
		}
		page, err := pageModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Title != "-isnil-" && updateData.Title != page.Title {
			page.Title = updateData.Title
		}
		if updateData.Image != "-isnil-" && updateData.Image != page.Image {
			page.Image = updateData.Image
		}
		if updateData.CategoryId != core.ToInt("-1") && updateData.CategoryId != page.CategoryId {
			page.CategoryId = updateData.CategoryId
		}
		if updateData.Content != "-isnil-" && updateData.Content != page.Content {
			page.Content = updateData.Content
		}
		err2 := page.UpdatePage()
		if err2 != nil {
			core.Error(c, 1, err2.Error())
			return
		}
		core.Success(c, 0, gin.H{
			"update": "true",
		})
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func DeletePage(c *gin.Context) {
	id, _ := c.Params.Get("page_id")
	if core.ToInt(id) > 0 {
		err := pageModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
