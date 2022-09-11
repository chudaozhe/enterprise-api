package admin

import (
	shortcutModel "enterprise-api/app/models/shortcut"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListShortcut(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "100")
	list, err := shortcutModel.List(false, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}

func CreateShortcut(c *gin.Context) {
	title := c.DefaultPostForm("title", "")
	if len(title) > 0 {
		shortcut := shortcutModel.Shortcut{
			Title:  title,
			Image:  c.DefaultPostForm("image", ""),               //图片
			Type:   core.ToInt(c.DefaultPostForm("type", "")),    //1普通分类2单页
			Url:    c.DefaultPostForm("url", ""),                 //链接
			Sort:   core.ToInt(c.DefaultPostForm("sort", "")),    //排序
			Status: core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		id, err := shortcut.CreateShortcut()
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
func DetailShortcut(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		detail, err := shortcutModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, detail)
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}

func ChangeShortcut(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		updateData := shortcutModel.Shortcut{
			Title:  c.DefaultPostForm("title", "-isnil-"),
			Image:  c.DefaultPostForm("image", "-isnil-"),        //图片
			Type:   core.ToInt(c.DefaultPostForm("type", "-1")),  //1普通分类2单页
			Url:    c.DefaultPostForm("url", "-isnil-"),          //链接
			Sort:   core.ToInt(c.DefaultPostForm("sort", "-1")),  //排序
			Status: core.ToInt(c.DefaultPostForm("status", "1")), //是否显示 1是 0否
		}
		shortcut, err := shortcutModel.FindById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, "记录不存在")
			return
		}
		if updateData.Title != "-isnil-" && updateData.Title != shortcut.Title {
			shortcut.Title = updateData.Title
		}
		if updateData.Image != "-isnil-" && updateData.Image != shortcut.Image {
			shortcut.Image = updateData.Image
		}
		if updateData.Type != core.ToInt("-1") && updateData.Type != shortcut.Type {
			shortcut.Type = updateData.Type
		}
		if updateData.Url != "-isnil-" && updateData.Url != shortcut.Url {
			shortcut.Url = updateData.Url
		}
		if updateData.Sort != core.ToInt("-1") && updateData.Sort != shortcut.Sort {
			shortcut.Sort = updateData.Sort
		}
		if updateData.Status != shortcut.Status {
			shortcut.Status = updateData.Status
		}
		err2 := shortcut.UpdateShortcut()
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

func DisplayShortcut(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		shortcut := &shortcutModel.Shortcut{Id: core.ToInt(id)}
		err := shortcut.ChangeState(1)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func HiddenShortcut(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		shortcut := &shortcutModel.Shortcut{Id: core.ToInt(id)}
		err := shortcut.ChangeState(0)
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"display": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
func DeleteShortcut(c *gin.Context) {
	id, _ := c.Params.Get("shortcut_id")
	if core.ToInt(id) > 0 {
		err := shortcutModel.DeleteById(core.ToInt(id))
		if err != nil {
			core.Error(c, 1, err.Error())
		} else {
			core.Success(c, 0, gin.H{"delete": true})
		}
	} else {
		core.Error(c, 1, "参数错误")
	}
}
