package user

import (
	articleModel "enterprise-api/app/models/article"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "10")
	keyword := c.DefaultQuery("keyword", "")
	categoryIds := []int{}
	if len(categoryId) > 0 && categoryId != "0" {
		categoryIds = append(categoryIds, core.ToInt(categoryId))
	}
	count, articles, err := articleModel.List(categoryIds, keyword, false, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": articles})
	}
}
func DetailArticle(c *gin.Context) {
	id, _ := c.Params.Get("article_id")
	if core.ToInt(id) > 0 {
		article, err := articleModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, article)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
