package admin

import (
	groupModel "enterprise-api/app/models/category"
	filesModel "enterprise-api/app/models/files"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListGroup(c *gin.Context) {
	list, err := groupModel.List(1)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, list)
	}
}

func CreateGroup(c *gin.Context) {
	var createGroupIn schemas.CreateGroupIn
	if err := c.ShouldBind(&createGroupIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	group := groupModel.Category{
		Name: createGroupIn.Name,
		Sort: createGroupIn.Sort, //排序
	}
	id, err := group.CreateFilegroup()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}

func DetailGroup(c *gin.Context) {
	var currentGroup schemas.CurrentGroup
	if err := c.ShouldBindUri(&currentGroup); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := groupModel.FindById(currentGroup.GroupId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, detail)
	}
}

func ChangeGroup(c *gin.Context) {
	var currentGroup schemas.CurrentGroup
	if err := c.ShouldBindUri(&currentGroup); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeGroupIn schemas.ChangeGroupIn
	if err := c.ShouldBind(&changeGroupIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	group, err := groupModel.FindById(currentGroup.GroupId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}
	group.Name = changeGroupIn.Name
	group.Sort = changeGroupIn.Sort
	err2 := group.UpdateFilegroup()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DeleteGroup(c *gin.Context) {
	var currentGroup schemas.CurrentGroup
	if err := c.ShouldBindUri(&currentGroup); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := groupModel.DeleteById(currentGroup.GroupId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		_ = filesModel.Updatefile2(currentGroup.GroupId) //把组下面的文件所属组改为0
		core.Success(c, 0, gin.H{"delete": true})
	}
}
