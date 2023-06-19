package admin

import (
	"enterprise-api/app/models"
	filesModel "enterprise-api/app/models/files"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
	"strings"
)

func ListFile(c *gin.Context) {
	var currentGroup schemas.CurrentGroup
	if err := c.ShouldBindUri(&currentGroup); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	var listFileIn schemas.ListFileIn
	if err := c.ShouldBindQuery(&listFileIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}

	count, list, err := filesModel.List(currentGroup.GroupId, listFileIn.Page, listFileIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"count": count, "list": list})
}

func CreateFile(c *gin.Context) {
	var createFileHeaderIn schemas.CreateFileHeaderIn
	if err := c.ShouldBindUri(&createFileHeaderIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var createFileIn schemas.CreateFileIn
	if err := c.ShouldBind(&createFileIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	file := filesModel.File{
		AdminId:    createFileHeaderIn.AdminId,
		CategoryId: createFileHeaderIn.GroupId,
		Url:        models.SavePhoto(createFileIn.Content),
		Title:      createFileIn.Title,
		Type:       createFileIn.Type,
		Size:       createFileIn.Size,
		Width:      createFileIn.Width,
		Height:     createFileIn.Height,
	}
	id, err := file.Createfile()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}

func DetailFile(c *gin.Context) {
	var currentFile schemas.CurrentFile
	if err := c.ShouldBindUri(&currentFile); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := filesModel.FindById(currentFile.FileId)
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, detail)
}

func ChangeFile(c *gin.Context) {
	var currentGroup schemas.CurrentGroup
	if err := c.ShouldBindUri(&currentGroup); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeFileIn schemas.ChangeFileIn
	if err := c.ShouldBind(&changeFileIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	fileIds := []int{}
	arr := strings.Split(changeFileIn.FileIds, ",")
	if len(arr) > 0 {
		for _, fileId := range arr {
			fid := core.ToInt(fileId)
			if fid == 0 {
				continue
			}
			fileIds = append(fileIds, fid)
		}
	}
	err := filesModel.Updatefile(fileIds, currentGroup.GroupId)
	if err != nil {
		core.Error(c, 1, "修改失败")
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DeleteFile(c *gin.Context) {
	var currentFile schemas.CurrentFile
	if err := c.ShouldBindUri(&currentFile); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err2 := filesModel.FindById(currentFile.FileId)
	if err2 != nil {
		core.Success(c, 0, gin.H{"delete": true})
		return
	}
	if len(detail.Url) > 0 && models.DeleteFile(detail.Url) {
		err := filesModel.DeleteById(currentFile.FileId)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "删除失败")
	}
}
