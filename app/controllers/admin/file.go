package admin

import (
	"enterprise-api/app/models"
	filesModel "enterprise-api/app/models/files"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
	"strings"
)

func ListFile(c *gin.Context) {
	groupId, _ := c.Params.Get("group_id")
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "20")
	count, list, err := filesModel.List(core.ToInt(groupId), core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"count": count, "list": list})
	}
}
func CreateFile(c *gin.Context) {
	adminId, _ := c.Params.Get("admin_id")
	groupId, _ := c.Params.Get("group_id")
	content := c.DefaultPostForm("content", "")
	if len(content) > 0 {
		file := filesModel.File{
			AdminId:    core.ToInt(adminId),
			CategoryId: core.ToInt(groupId),
			Url:        models.SavePhoto(content),
			Title:      c.DefaultPostForm("title", ""),
			Type:       c.DefaultPostForm("type", ""),
			Size:       core.ToInt(c.DefaultPostForm("size", "")),
			Width:      core.ToInt(c.DefaultPostForm("width", "")),
			Height:     core.ToInt(c.DefaultPostForm("height", "")),
		}
		id, err := file.Createfile()
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
func DetailFile(c *gin.Context) {
	id, _ := c.Params.Get("file_id")
	if core.ToInt(id) > 0 {
		detail, err := filesModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeFile(c *gin.Context) {
	id, _ := c.Params.Get("group_id")
	fileIdStr := c.DefaultPostForm("file_ids", "") //
	fileIds := []int{}
	if len(fileIdStr) > 0 {
		arr := strings.Split(fileIdStr, ",")
		if len(arr) > 0 { //数组转切片
			for _, fileId := range arr {
				if fileId == "0" {
					continue
				}
				fileIds = append(fileIds, core.ToInt(fileId))
			}
		}
		err := filesModel.Updatefile(fileIds, core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "修改失败")
		} else {
			core.Success(c, 0, gin.H{"update": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func DeleteFile(c *gin.Context) {
	id, _ := c.Params.Get("file_id")
	if core.ToInt(id) > 0 {
		detail, err2 := filesModel.FindById(core.ToInt(id))
		if err2 != nil {
			core.Success(c, 0, gin.H{"delete": true})
			return
		}
		if len(detail.Url) > 0 && models.DeleteFile(detail.Url) {
			err := filesModel.DeleteById(core.ToInt(id))
			if err != nil {
				core.Error(c, 1, err.Error())
			} else {
				core.Success(c, 0, gin.H{"delete": true})
			}
		} else {
			core.Error(c, 1, "删除失败")
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
