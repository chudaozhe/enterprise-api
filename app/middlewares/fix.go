package middlewares

import (
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
)

// 修复一些参数
func FixParam() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()

		page := params.Get("page")
		max := params.Get("max")
		if core.ToInt(page) < 1 {
			page = "1"
		}
		if core.ToInt(max) < 1 {
			max = "10"
		}

		// 修改参数
		params.Set("page", page)
		params.Set("max", max)

		// 将修改后的参数重新设置到请求中
		c.Request.URL.RawQuery = params.Encode()

		c.Next()
	}
}
