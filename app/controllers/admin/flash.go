package admin

import (
	flashModel "enterprise-api/app/models/flash"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListFlash(c *gin.Context) {
	var listFlashIn schemas.ListFlashIn
	if err := c.ShouldBindQuery(&listFlashIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	list, err := flashModel.List(false, listFlashIn.Page, listFlashIn.Max)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, list)
	}
}

func CreateFlash(c *gin.Context) {
	var createFlashIn schemas.CreateFlashIn
	if err := c.ShouldBind(&createFlashIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	page := flashModel.Flash{
		Image:  createFlashIn.Image,
		Title:  createFlashIn.Title,
		Url:    createFlashIn.Url,
		Sort:   createFlashIn.Sort,
		Status: createFlashIn.Status, //是否显示 1是 0否
	}
	id, err := page.CreateFlash()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}

func DetailFlash(c *gin.Context) {
	var currentFlash schemas.CurrentFlash
	if err := c.ShouldBindUri(&currentFlash); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := flashModel.FindById(currentFlash.FlashId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, detail)
	}
}

func ChangeFlash(c *gin.Context) {
	var currentFlash schemas.CurrentFlash
	if err := c.ShouldBindUri(&currentFlash); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeFlashIn schemas.ChangeFlashIn
	if err := c.ShouldBind(&changeFlashIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	flash, err := flashModel.FindById(currentFlash.FlashId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}
	flash.Image = changeFlashIn.Image
	flash.Title = changeFlashIn.Title
	flash.Url = changeFlashIn.Url
	flash.Sort = changeFlashIn.Sort
	flash.Status = changeFlashIn.Status //是否显示 1是 0否
	err2 := flash.UpdateFlash()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DisplayFlash(c *gin.Context) {
	var currentFlash schemas.CurrentFlash
	if err := c.ShouldBindUri(&currentFlash); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	flash := &flashModel.Flash{Id: currentFlash.FlashId}
	err := flash.ChangeState(1)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"display": true})
	}
}
func HiddenFlash(c *gin.Context) {
	var currentFlash schemas.CurrentFlash
	if err := c.ShouldBindUri(&currentFlash); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	flash := &flashModel.Flash{Id: currentFlash.FlashId}
	err := flash.ChangeState(0)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"hidden": true})
	}
}

func DeleteFlash(c *gin.Context) {
	var currentFlash schemas.CurrentFlash
	if err := c.ShouldBindUri(&currentFlash); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := flashModel.DeleteById(currentFlash.FlashId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"delete": true})
	}
}
