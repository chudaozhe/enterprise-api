package admin

import (
	casesModel "enterprise-api/app/models/cases"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCase(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var listCaseIn schemas.ListCaseIn
	if err := c.ShouldBindQuery(&listCaseIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	count, list, err := casesModel.List(currentCategory.CategoryId, listCaseIn.Keyword, false, listCaseIn.Page, listCaseIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"count": count, "list": list})
}

func CreateCase(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var createCaseIn schemas.CreateCaseIn
	if err := c.ShouldBind(&createCaseIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	cases := casesModel.Case{
		Title:       createCaseIn.Title,
		Content:     createCaseIn.Content,
		CategoryId:  currentCategory.CategoryId, //分类id
		Description: createCaseIn.Description,   //描述
		Image:       createCaseIn.Image,         //图片
		Images:      createCaseIn.Images,        //图集
		Sort:        createCaseIn.Sort,          //排序
		Url:         createCaseIn.Url,           //外部链接
		Status:      createCaseIn.Status,        //是否显示 1是 0否
	}
	id, err := cases.CreateCase()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}

func DetailCase(c *gin.Context) {
	var currentCase schemas.CurrentCase
	if err := c.ShouldBindUri(&currentCase); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := casesModel.FindById(currentCase.CaseId)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, detail)
}

func ChangeCase(c *gin.Context) {
	var changeCaseHeaderIn schemas.ChangeCaseHeaderIn
	if err := c.ShouldBindUri(&changeCaseHeaderIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	var changeCaseIn schemas.ChangeCaseIn
	if err := c.ShouldBind(&changeCaseIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	cases, err := casesModel.FindById(changeCaseHeaderIn.CaseId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}
	cases.Title = changeCaseIn.Title
	cases.Content = changeCaseIn.Content
	cases.CategoryId = changeCaseHeaderIn.CategoryId //分类id
	cases.Description = changeCaseIn.Description     //描述
	cases.Image = changeCaseIn.Image                 //图片
	cases.Images = changeCaseIn.Images               //图集
	cases.Sort = changeCaseIn.Sort                   //排序
	cases.Url = changeCaseIn.Url                     //外部链接
	cases.Status = changeCaseIn.Status               //是否显示 1是 0否
	err2 := cases.UpdateCase()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DisplayCase(c *gin.Context) {
	var currentCase schemas.CurrentCase
	if err := c.ShouldBindUri(&currentCase); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	cases := &casesModel.Case{Id: currentCase.CaseId}
	err := cases.ChangeState(1)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"display": true})
}

func HiddenCase(c *gin.Context) {
	var currentCase schemas.CurrentCase
	if err := c.ShouldBindUri(&currentCase); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	cases := &casesModel.Case{Id: currentCase.CaseId}
	err := cases.ChangeState(0)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"hidden": true})
}

func DeleteCase(c *gin.Context) {
	var currentCase schemas.CurrentCase
	if err := c.ShouldBindUri(&currentCase); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := casesModel.DeleteById(currentCase.CaseId)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"delete": true})
}
