package user

import (
	flashModel "enterprise-api/app/models/flash"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func ListFlash(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	max := c.DefaultQuery("max", "100")
	list, err := flashModel.List(true, core.ToInt(page), core.ToInt(max))
	if err != nil {
		core.Error(c, 1, err.Error())
	} else {
		core.Success(c, 0, list)
	}
}
