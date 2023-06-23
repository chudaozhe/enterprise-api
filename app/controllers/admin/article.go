package admin

import (
	articleModel "enterprise-api/app/models/article"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
	"strings"
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
	categoryIds := []int{}
	if currentCategory.CategoryId > 0 {
		categoryIds = append(categoryIds, currentCategory.CategoryId)
	}
	count, articles, err := articleModel.List(categoryIds, listArticleIn.Keyword, false, listArticleIn.Page, listArticleIn.Max)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": articles})
	}
}

func List2Article(c *gin.Context) {
	var list2ArticleIn schemas.List2ArticleIn
	if err := c.ShouldBindQuery(&list2ArticleIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	category_ids := []int{}
	if len(list2ArticleIn.CategoryIds) > 0 {
		arr := strings.Split(list2ArticleIn.CategoryIds, ",")
		if len(arr) > 0 {
			for _, categoryId := range arr {
				cid := core.ToInt(categoryId)
				if cid == 0 {
					continue
				}
				category_ids = append(category_ids, cid)
			}
		}
	}
	count, articles, err := articleModel.List(category_ids, list2ArticleIn.Keyword, false, list2ArticleIn.Page, list2ArticleIn.Max)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": articles})
	}
}

func CreateArticle(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var createArticleIn schemas.CreateArticleIn
	if err := c.ShouldBind(&createArticleIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	article := articleModel.Article{
		Title:       createArticleIn.Title,
		Content:     createArticleIn.Content,
		CategoryId:  currentCategory.CategoryId,  //分类id
		Keywords:    createArticleIn.Keywords,    //关键词
		Description: createArticleIn.Description, //描述
		Image:       createArticleIn.Image,       //图片
		Images:      createArticleIn.Images,      //图集
		Sort:        createArticleIn.Sort,        //排序
		Source:      createArticleIn.Source,      //来源
		Author:      createArticleIn.Author,      //作者
		Url:         createArticleIn.Url,         //外部链接
		Status:      createArticleIn.Status,      //是否显示 1是 0否
	}
	id, err := article.CreateArticle()
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"id": id})
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
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, article)
	}
}

func ChangeArticle(c *gin.Context) {
	var changeArticleHeaderIn schemas.ChangeArticleHeaderIn
	if err := c.ShouldBindUri(&changeArticleHeaderIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	var changeArticleIn schemas.ChangeArticleIn
	if err := c.ShouldBind(&changeArticleIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	article, err := articleModel.FindById(changeArticleHeaderIn.ArticleId)
	if err != nil {
		_ = c.Error(err)
		return
	}
	article.Title = changeArticleIn.Title
	article.Content = changeArticleIn.Content
	article.CategoryId = changeArticleHeaderIn.CategoryId //分类id
	article.Keywords = changeArticleIn.Keywords           //关键词
	article.Description = changeArticleIn.Description     //描述
	article.Image = changeArticleIn.Image                 //图片
	article.Images = changeArticleIn.Images               //图集
	article.Sort = changeArticleIn.Sort                   //排序
	article.Source = changeArticleIn.Source               //来源
	article.Author = changeArticleIn.Author               //作者
	article.Url = changeArticleIn.Url                     //外部链接
	article.Status = changeArticleIn.Status               //是否显示 1是 0否

	err2 := article.UpdateArticle()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	} else {
		core.Success(c, 0, gin.H{"update": true})
	}
}

func DisplayArticle(c *gin.Context) {
	var currentArticle schemas.CurrentArticle
	if err := c.ShouldBindUri(&currentArticle); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	article := &articleModel.Article{Id: currentArticle.ArticleId}
	err := article.ChangeArticleState(1)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"display": true})
	}
}

func HiddenArticle(c *gin.Context) {
	var currentArticle schemas.CurrentArticle
	if err := c.ShouldBindUri(&currentArticle); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	article := &articleModel.Article{Id: currentArticle.ArticleId}
	err := article.ChangeArticleState(0)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"hidden": true})
	}
}

func DeleteArticle(c *gin.Context) {
	var currentArticle schemas.CurrentArticle
	if err := c.ShouldBindUri(&currentArticle); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := articleModel.DeleteArticleById(currentArticle.ArticleId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"delete": true})
	}
}
