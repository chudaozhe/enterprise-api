package user

import (
	shortcutModel "enterprise-api/app/models/shortcut"
	"enterprise-api/app/schemas"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListShortcut(c *gin.Context) {
	var listShortcutIn schemas.ListShortcutIn
	if err := c.ShouldBindQuery(&listShortcutIn); err != nil {
		_ = c.Error(err)
		return
	}
	list, err := shortcutModel.List(true, listShortcutIn.Page, listShortcutIn.Max)
	if err != nil {
		_ = c.Error(err)
		return
	} else {
		core.Success(c, 0, list)
	}
}
