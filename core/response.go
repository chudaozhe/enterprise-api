package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 正确状态处理
func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"err":  code,
		"data": data,
	})
}

// 错误状态处理
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusBadRequest,
		gin.H{
			"err": code,
			"msg": msg,
		})
}
