package user

import (
	articleModel "enterprise-api/app/models/article"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListArticle(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var listArticleIn schemas.ListArticleIn
	if err := c.ShouldBindQuery(&listArticleIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var categoryIds []int
	if currentCategory.CategoryId > 0 {
		categoryIds = append(categoryIds, currentCategory.CategoryId)
	}
	count, articles, err := articleModel.List(categoryIds, listArticleIn.Keyword, false, listArticleIn.Page, listArticleIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": articles})
	}
}
func DetailArticle(c *gin.Context) {
	var detailArticleIn schemas.DetailArticleIn
	if err := c.ShouldBindUri(&detailArticleIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	article, err := articleModel.FindById(detailArticleIn.ArticleId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, article)
	}
}
