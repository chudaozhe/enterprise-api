package admin

import (
	groupModel "enterprise-api/app/models/category"
	filesModel "enterprise-api/app/models/files"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListGroup(c *gin.Context) {
	list, err := groupModel.List(1)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
func CreateGroup(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if len(name) > 0 {
		group := groupModel.Category{
			Name: name,
			Sort: core.ToInt(c.DefaultPostForm("sort", "")), //排序
		}
		id, err := group.CreateFilegroup()
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
func DetailGroup(c *gin.Context) {
	id, _ := c.Params.Get("group_id")
	if core.ToInt(id) > 0 {
		detail, err := groupModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeGroup(c *gin.Context) {
	id, _ := c.Params.Get("group_id")
	if core.ToInt(id) > 0 {
		updateData := groupModel.Category{
			Name: c.DefaultPostForm("name", "-isnil-"),
			Sort: core.ToInt(c.DefaultPostForm("sort", "-1")), //排序
		}
		group, err := groupModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Name != "-isnil-" && updateData.Name != group.Name {
			group.Name = updateData.Name
		}
		if updateData.Sort != core.ToInt("-1") && updateData.Sort != group.Sort {
			group.Sort = updateData.Sort
		}

		err2 := group.UpdateFilegroup()
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

func DeleteGroup(c *gin.Context) {
	id, _ := c.Params.Get("group_id")
	if core.ToInt(id) > 0 {
		err := groupModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			filesModel.Updatefile2(core.ToInt(id)) //把组下面的文件所属组改为0
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
