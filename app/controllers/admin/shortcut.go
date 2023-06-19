package admin

import (
	shortcutModel "enterprise-api/app/models/shortcut"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListShortcut(c *gin.Context) {
	var listShortcutIn schemas.ListShortcutIn
	if err := c.ShouldBind(&listShortcutIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	list, err := shortcutModel.List(false, listShortcutIn.Page, listShortcutIn.Max)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}

func CreateShortcut(c *gin.Context) {
	var createShortcutIn schemas.CreateShortcutIn
	if err := c.ShouldBind(&createShortcutIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	shortcut := shortcutModel.Shortcut{
		Title:  createShortcutIn.Title,
		Image:  createShortcutIn.Image,  //图片
		Type:   createShortcutIn.Type,   //1普通分类2单页
		Url:    createShortcutIn.Url,    //链接
		Sort:   createShortcutIn.Sort,   //排序
		Status: createShortcutIn.Status, //是否显示 1是 0否
	}
	id, err := shortcut.CreateShortcut()
	if err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	core.Success(c, 0, gin.H{"id": id})
}
func DetailShortcut(c *gin.Context) {
	var currentShortcut schemas.CurrentShortcut
	if err := c.ShouldBindUri(&currentShortcut); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	detail, err := shortcutModel.FindById(currentShortcut.ShortcutId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, detail)
	}
}

func ChangeShortcut(c *gin.Context) {
	var currentShortcut schemas.CurrentShortcut
	if err := c.ShouldBindUri(&currentShortcut); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	var changeShortcutIn schemas.ChangeShortcutIn
	if err := c.ShouldBind(&changeShortcutIn); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	shortcut, err := shortcutModel.FindById(currentShortcut.ShortcutId)
	if err != nil {
		core.Error(c, 1, "记录不存在")
		return
	}
	shortcut.Title = changeShortcutIn.Title
	shortcut.Image = changeShortcutIn.Image   //图片
	shortcut.Type = changeShortcutIn.Type     //1普通分类2单页
	shortcut.Url = changeShortcutIn.Url       //链接
	shortcut.Sort = changeShortcutIn.Sort     //排序
	shortcut.Status = changeShortcutIn.Status //是否显示 1是 0否
	err2 := shortcut.UpdateShortcut()
	if err2 != nil {
		core.Error(c, 1, err2.Error())
		return
	}
	core.Success(c, 0, gin.H{"update": true})
}

func DisplayShortcut(c *gin.Context) {
	var currentShortcut schemas.CurrentShortcut
	if err := c.ShouldBindUri(&currentShortcut); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	shortcut := &shortcutModel.Shortcut{Id: currentShortcut.ShortcutId}
	err := shortcut.ChangeState(1)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"display": true})
	}
}

func HiddenShortcut(c *gin.Context) {
	var currentShortcut schemas.CurrentShortcut
	if err := c.ShouldBindUri(&currentShortcut); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	shortcut := &shortcutModel.Shortcut{Id: currentShortcut.ShortcutId}
	err := shortcut.ChangeState(0)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"hidden": true})
	}
}

func DeleteShortcut(c *gin.Context) {
	var currentShortcut schemas.CurrentShortcut
	if err := c.ShouldBindUri(&currentShortcut); err != nil {
		core.Error(c, 1, err.Error())
		return
	}
	err := shortcutModel.DeleteById(currentShortcut.ShortcutId)
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, gin.H{"delete": true})
	}
}
