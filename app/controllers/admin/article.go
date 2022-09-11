package admin

import (
	articleModel "enterprise-api/app/models/article"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
	"strings"
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
func List2Article(c *gin.Context) {
	categoryIds := c.DefaultQuery("category_ids", "")
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "10")
	keyword := c.DefaultQuery("keyword", "")
	category_ids := []int{}
	if len(categoryIds) > 0 {
		arr := strings.Split(categoryIds, ",")
		if len(arr) > 0 { //数组转切片
			for _, categoryId := range arr {
				if categoryId == "0" {
					continue
				}
				category_ids = append(category_ids, core.ToInt(categoryId))
			}
		}
	}
	count, articles, err := articleModel.List(category_ids, keyword, false, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": articles})
	}
}

func CreateArticle(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	if len(title) > 0 && len(content) > 0 {
		article := articleModel.Article{
			Title:       title,
			Content:     content,
			CategoryId:  core.ToInt(categoryId),                       //分类id
			Keywords:    c.DefaultPostForm("keywords", ""),            //关键词
			Description: c.DefaultPostForm("description", ""),         //描述
			Image:       c.DefaultPostForm("image", ""),               //图片
			Images:      c.DefaultPostForm("images", ""),              //图集
			Sort:        core.ToInt(c.DefaultPostForm("sort", "")),    //排序
			Source:      c.DefaultPostForm("source", ""),              //来源
			Author:      c.DefaultPostForm("author", ""),              //作者
			Url:         c.DefaultPostForm("url", ""),                 //外部链接
			Status:      core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		id, err := article.CreateArticle()
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

func ChangeArticle(c *gin.Context) {
	id, _ := c.Params.Get("article_id")
	categoryId, _ := c.Params.Get("category_id")
	if core.ToInt(id) > 0 {
		updateData := articleModel.Article{
			Title:       c.DefaultPostForm("title", "-isnil-"),
			Content:     c.DefaultPostForm("content", "-isnil-"),
			CategoryId:  core.ToInt(categoryId),                      //分类id
			Keywords:    c.DefaultPostForm("keywords", "-isnil-"),    //关键词
			Description: c.DefaultPostForm("description", "-isnil-"), //描述
			Image:       c.DefaultPostForm("image", "-isnil-"),       //图片
			Images:      c.DefaultPostForm("images", "-isnil-"),      //图集
			Sort:        core.ToInt(c.DefaultPostForm("sort", "")),   //排序
			Source:      c.DefaultPostForm("source", "-isnil-"),      //来源
			Author:      c.DefaultPostForm("author", "-isnil-"),      //作者
			Url:         c.DefaultPostForm("url", "-isnil-"),         //外部链接
			Status:      core.ToInt(c.DefaultPostForm("status", "")), //是否显示 1是 0否
		}
		article, err := articleModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Title != "-isnil-" && updateData.Title != article.Title {
			article.Title = updateData.Title
		}
		if updateData.Content != "-isnil-" && updateData.Content != article.Content {
			article.Content = updateData.Content
		}
		if updateData.CategoryId != article.CategoryId {
			article.CategoryId = updateData.CategoryId
		}
		if updateData.Keywords != "-isnil-" && updateData.Keywords != article.Keywords {
			article.Keywords = updateData.Keywords
		}
		if updateData.Description != "-isnil-" && updateData.Description != article.Description {
			article.Description = updateData.Description
		}
		if updateData.Image != "-isnil-" && updateData.Image != article.Image {
			article.Image = updateData.Image
		}
		if updateData.Images != "-isnil-" && updateData.Images != article.Images {
			article.Images = updateData.Images
		}
		if updateData.Sort != article.Sort {
			article.Sort = updateData.Sort
		}
		if updateData.Source != "-isnil-" && updateData.Source != article.Source {
			article.Source = updateData.Source
		}
		if updateData.Author != "-isnil-" && updateData.Author != article.Author {
			article.Author = updateData.Author
		}
		if updateData.Url != "-isnil-" && updateData.Url != article.Url {
			article.Url = updateData.Url
		}
		if updateData.Status != article.Status {
			article.Status = updateData.Status
		}
		err2 := article.UpdateArticle()
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

func DisplayArticle(c *gin.Context) {
	id, _ := c.Params.Get("article_id")
	if core.ToInt(id) > 0 {
		article := &articleModel.Article{Id: core.ToInt(id)}
		err := article.ChangeArticleState(1)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func HiddenArticle(c *gin.Context) {
	id, _ := c.Params.Get("article_id")
	if core.ToInt(id) > 0 {
		article := &articleModel.Article{Id: core.ToInt(id)}
		err := article.ChangeArticleState(0)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"hidden": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func DeleteArticle(c *gin.Context) {
	id, _ := c.Params.Get("article_id")
	if core.ToInt(id) > 0 {
		err := articleModel.DeleteArticleById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
