package user

import (
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func CreateTest(c *gin.Context) {
	core.Success(c, 0, "ok")
}
