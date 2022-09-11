package admin

import (
	flashModel "enterprise-api/app/models/flash"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListFlash(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "100")
	list, err := flashModel.List(false, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}

func CreateFlash(c *gin.Context) {
	image := c.DefaultPostForm("image", "")
	if len(image) > 0 {
		page := flashModel.Flash{
			Image:  image,
			Title:  c.DefaultPostForm("title", ""),
			Url:    c.DefaultPostForm("url", ""),
			Sort:   core.ToInt(c.DefaultPostForm("sort", "")),
			Status: core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		id, err := page.CreateFlash()
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
func DetailFlash(c *gin.Context) {
	id, _ := c.Params.Get("flash_id")
	if core.ToInt(id) > 0 {
		detail, err := flashModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeFlash(c *gin.Context) {
	id, _ := c.Params.Get("flash_id")
	if core.ToInt(id) > 0 {
		updateData := flashModel.Flash{
			Image:  c.DefaultPostForm("image", "-isnil-"),
			Title:  c.DefaultPostForm("title", "-isnil-"),
			Url:    c.DefaultPostForm("url", "-isnil-"),
			Sort:   core.ToInt(c.DefaultPostForm("sort", "-1")),
			Status: core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		flash, err := flashModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Image != "-isnil-" && updateData.Image != flash.Image {
			flash.Image = updateData.Image
		}
		if updateData.Title != "-isnil-" && updateData.Title != flash.Title {
			flash.Title = updateData.Title
		}
		if updateData.Url != "-isnil-" && updateData.Url != flash.Url {
			flash.Url = updateData.Url
		}
		if updateData.Sort != -1 && updateData.Sort != flash.Sort {
			flash.Sort = updateData.Sort
		}
		if updateData.Status != flash.Status {
			flash.Status = updateData.Status
		}
		err2 := flash.UpdateFlash()
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

func DisplayFlash(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		flash := &flashModel.Flash{Id: core.ToInt(id)}
		err := flash.ChangeState(1)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func HiddenFlash(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		flash := &flashModel.Flash{Id: core.ToInt(id)}
		err := flash.ChangeState(0)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func DeleteFlash(c *gin.Context) {
	id, _ := c.Params.Get("flash_id")
	if core.ToInt(id) > 0 {
		err := flashModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
