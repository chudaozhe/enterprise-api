package user

import (
	shortcutModel "enterprise-api/app/models/shortcut"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListShortcut(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "100")
	list, err := shortcutModel.List(true, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
