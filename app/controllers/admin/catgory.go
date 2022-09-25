package admin

import (
	categoryModel "enterprise-api/app/models/category"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	type0 := c.DefaultQuery("type", "")
	list, err := categoryModel.List(core.ToInt(type0))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
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
	name := c.DefaultPostForm("name", "")
	type0 := core.ToInt(c.DefaultPostForm("type", "2"))
	if len(name) > 0 {
		_, err0 := categoryModel.FindByName(name, type0)
		if err0 != nil {
			category := categoryModel.Category{
				Name:     name,
				Memo:     c.DefaultPostForm("memo", ""),
				Type:     core.ToInt(c.DefaultPostForm("type", "1")),
				ParentId: core.ToInt(c.DefaultPostForm("parent_id", "")),
				Sort:     core.ToInt(c.DefaultPostForm("sort", "")),
			}
			id, err := category.CreateCategory()
			if err != nil {
				core.Error(c, 1, err.Error())
				return
			}
			core.Success(c, 0, gin.H{
				"id": id,
			})
		} else {
			core.Error(c, 405, "记录已存在")
		}
	} else {
		core.Success(c, 400, "参数错误")
	}
}
func DetailCategory(c *gin.Context) {
	id, _ := c.Params.Get("category_id")
	if core.ToInt(id) > 0 {
		detail, err := categoryModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeCategory(c *gin.Context) {
	id, _ := c.Params.Get("category_id")
	if core.ToInt(id) > 0 {
		updateData := categoryModel.Category{
			Name:     c.DefaultPostForm("name", "-isnil-"),
			Memo:     c.DefaultPostForm("memo", "-isnil-"),
			Type:     core.ToInt(c.DefaultPostForm("type", "-1")),
			ParentId: core.ToInt(c.DefaultPostForm("parent_id", "-1")),
			Sort:     core.ToInt(c.DefaultPostForm("sort", "-1")),
		}
		category, err := categoryModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Name != "-isnil-" && updateData.Name != category.Name {
			category.Name = updateData.Name
		}
		if updateData.Memo != "-isnil-" && updateData.Memo != category.Memo {
			category.Memo = updateData.Memo
		}
		if updateData.Type != core.ToInt("-1") && updateData.Type != category.Type {
			category.Type = updateData.Type
		}
		if updateData.ParentId != core.ToInt("-1") && updateData.ParentId != category.ParentId {
			category.ParentId = updateData.ParentId
		}
		if updateData.Sort != core.ToInt("-1") && updateData.Sort != category.Sort {
			category.Sort = updateData.Sort
		}
		err2 := category.UpdateCategory()
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

func DeleteCategory(c *gin.Context) {
	id, _ := c.Params.Get("category_id")
	if core.ToInt(id) > 0 {
		err := categoryModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
