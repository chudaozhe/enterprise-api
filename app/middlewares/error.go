package middlewares

import (
	"enterprise-api/app/models/errors"
	"enterprise-api/core"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// Error 拦截并处理 c.Errors 中的错误
func Error() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先调用c.Next()执行后面的中间件
		// 所有中间件及router处理完毕后从这里开始执行
		// 检查c.Errors中是否有错误
		for _, errorItem := range c.Errors {
			err := errorItem.Err
			//fmt.Printf("Type: %T\n", err)
			switch err.(type) {
			case validator.ValidationErrors:
				c.JSON(http.StatusOK, gin.H{
					"err": 400,
					"msg": core.Translate(err),
				})
			case *errors.CustomError:
				// 若是自定义的错误则将err、msg返回
				if customErr, ok := err.(*errors.CustomError); ok {
					c.JSON(http.StatusOK, gin.H{
						"err": customErr.Err,
						"msg": customErr.Msg,
					})
				}
			default:
				// 非自定义错误则返回详细错误信息err.Error()
				c.JSON(http.StatusOK, gin.H{
					"err": 500,
					"msg": err.Error(), //服务器异常
				})
			}
			return // 检查一个错误就行
		}
	}
}
