package middlewares

import (
	"enterprise-api/app/models"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

func JWTAuth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		signToken := c.Request.Header.Get("Authorization")
		roleId, ok := c.Params.Get(role + "_id")
		if signToken == "" || !ok { //未传递user_id
			core.Error(c, 400, "无效的id")
			c.Abort()
			return
		}
		if role == "admin" || (role == "user" && core.ToInt(roleId) > 0) {
			myclaims, err := models.VerifyToken(signToken)
			if err != nil {
				core.Error(c, 401, "token校验失败")
				c.Abort()
				return
			}
			//c.Set("userid", myclaims.Id)
			if myclaims.Id == 0 || myclaims.Id != core.ToInt(roleId) {
				core.Error(c, 401, "token校验失败")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
