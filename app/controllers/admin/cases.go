package admin

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

func CreateCase(c *gin.Context) {
	categoryId, _ := c.Params.Get("category_id")
	title := c.DefaultPostForm("title", "")
	content := c.DefaultPostForm("content", "")
	if len(title) > 0 && len(content) > 0 {
		cases := casesModel.Case{
			Title:       title,
			Content:     content,
			CategoryId:  core.ToInt(categoryId),                       //分类id
			Description: c.DefaultPostForm("description", ""),         //描述
			Image:       c.DefaultPostForm("image", ""),               //图片
			Images:      c.DefaultPostForm("images", ""),              //图集
			Sort:        core.ToInt(c.DefaultPostForm("sort", "")),    //排序
			Url:         c.DefaultPostForm("url", ""),                 //外部链接
			Status:      core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		id, err := cases.CreateCase()
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

func ChangeCase(c *gin.Context) {
	id, _ := c.Params.Get("case_id")
	categoryId, _ := c.Params.Get("category_id")
	if core.ToInt(id) > 0 {
		updateData := casesModel.Case{
			Title:       c.DefaultPostForm("title", ""),
			Content:     c.DefaultPostForm("keywords", ""),
			CategoryId:  core.ToInt(categoryId),                      //分类id
			Description: c.DefaultPostForm("description", ""),        //描述
			Image:       c.DefaultPostForm("image", ""),              //图片
			Images:      c.DefaultPostForm("images", ""),             //图集
			Sort:        core.ToInt(c.DefaultPostForm("sort", "")),   //排序
			Url:         c.DefaultPostForm("url", ""),                //外部链接
			Status:      core.ToInt(c.DefaultPostForm("status", "")), //是否显示 1是 0否
		}
		cases, err := casesModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Title != "-isnil-" && updateData.Title != cases.Title {
			cases.Title = updateData.Title
		}
		if updateData.Content != "-isnil-" && updateData.Content != cases.Content {
			cases.Content = updateData.Content
		}
		if updateData.CategoryId != cases.CategoryId {
			cases.CategoryId = updateData.CategoryId
		}
		if updateData.Description != "-isnil-" && updateData.Description != cases.Description {
			cases.Description = updateData.Description
		}
		if updateData.Image != "-isnil-" && updateData.Image != cases.Image {
			cases.Image = updateData.Image
		}
		if updateData.Images != "-isnil-" && updateData.Images != cases.Images {
			cases.Images = updateData.Images
		}
		if updateData.Sort != cases.Sort {
			cases.Sort = updateData.Sort
		}
		if updateData.Url != "-isnil-" && updateData.Url != cases.Url {
			cases.Url = updateData.Url
		}
		if updateData.Status != cases.Status {
			cases.Status = updateData.Status
		}
		err2 := cases.UpdateCase()
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

func DisplayCase(c *gin.Context) {
	id, _ := c.Params.Get("case_id")
	if core.ToInt(id) > 0 {
		cases := &casesModel.Case{Id: core.ToInt(id)}
		err := cases.ChangeState(1)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func HiddenCase(c *gin.Context) {
	id, _ := c.Params.Get("case_id")
	if core.ToInt(id) > 0 {
		cases := &casesModel.Case{Id: core.ToInt(id)}
		err := cases.ChangeState(0)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func DeleteCase(c *gin.Context) {
	id, _ := c.Params.Get("case_id")
	if core.ToInt(id) > 0 {
		err := casesModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
