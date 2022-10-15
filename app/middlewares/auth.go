package middlewares

import (
	"enterprise-api/app/models"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

// api鉴权中间件
func Auth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		roleId, ok := c.Params.Get(role + "_id")
		if !ok { //未传递user_id
			core.Error(c, 400, "无效的id")
			c.Abort()
			return
		}
		if role == "admin" || (role == "user" && core.ToInt(roleId) > 0) {
			err := models.CheckToken(core.ToInt(roleId), token, role)
			if err != nil {
				core.Error(c, 401, err.Error())
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
