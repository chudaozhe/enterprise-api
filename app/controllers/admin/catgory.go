package admin

import (
	categoryModel "enterprise-api/app/models/category"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var listCategoryIn schemas.ListCategoryIn
	if err := c.ShouldBindQuery(&listCategoryIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	list, err := categoryModel.List(listCategoryIn.Type)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, list)
}

func List2Category(c *gin.Context) {
	var boxArr []map[string]interface{} //最外层map切片
	boxArr = append(boxArr, map[string]interface{}{"label": "新闻资讯", "value": 2})
	boxArr = append(boxArr, map[string]interface{}{"label": "技术支持", "value": 5})

	var type2arr []map[string]interface{}
	var type5arr []map[string]interface{}

	list, _ := categoryModel.List2()
	for _, item := range list {
		obj := make(map[string]interface{})
		if item.Type == 2 {
			obj["label"] = item.Name
			obj["value"] = item.Id
			type2arr = append(type2arr, obj)
		} else if item.Type == 5 {
			obj["label"] = item.Name
			obj["value"] = item.Id
			type5arr = append(type5arr, obj)
		}
	}
	boxArr[0]["children"] = type2arr
	boxArr[1]["children"] = type5arr
	core.Success(c, 0, boxArr)
}

func CreateCategory(c *gin.Context) {
	var createCategoryIn schemas.CreateCategoryIn
	if err := c.ShouldBind(&createCategoryIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	_, err0 := categoryModel.FindByName(createCategoryIn.Name, createCategoryIn.Type)
	if err0 != nil {
		category := categoryModel.Category{
			Name:     createCategoryIn.Name,
			Memo:     createCategoryIn.Memo,
			Type:     createCategoryIn.Type, //类型：1文件 2新闻 3技术支持
			ParentId: createCategoryIn.ParentId,
			Sort:     createCategoryIn.Sort,
		}
		id, err := category.CreateCategory()
		if err != nil {
			core.Error(c, 1, err.Error())
			return
		}
		core.Success(c, 0, gin.H{"id": id})
	} else {
		core.Error(c, 405, "记录已存在")
	}
}

func DetailCategory(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := categoryModel.FindById(currentCategory.CategoryId)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, detail)
}

func ChangeCategory(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeCategoryIn schemas.ChangeCategoryIn
	if err := c.ShouldBind(&changeCategoryIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	category, err := categoryModel.FindById(currentCategory.CategoryId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}

	category.Name = changeCategoryIn.Name
	category.Memo = changeCategoryIn.Memo
	category.Type = changeCategoryIn.Type
	category.ParentId = changeCategoryIn.ParentId
	category.Sort = changeCategoryIn.Sort
	err2 := category.UpdateCategory()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DeleteCategory(c *gin.Context) {
	var currentCategory schemas.CurrentCategory
	if err := c.ShouldBindUri(&currentCategory); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := categoryModel.DeleteById(currentCategory.CategoryId)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"delete": true})
}
