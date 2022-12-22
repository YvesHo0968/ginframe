package middleware

import (
	"ginFrame/common"
	"ginFrame/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerException 全局异常处理
func HandlerException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				config.Log.Error().Interface("HandlerException", err).Msg("Handler Exception")
				common.ServerError(c, http.StatusInternalServerError, "Internal Server Error")
				c.Abort()
			}
		}()

		c.Next()

		// 处理逻辑收集的一般错误，不影响程序执行
		if length := len(c.Errors); length > 0 {
			for _, v := range c.Errors {
				config.Log.Warn().Interface("Errors", v).Msg("Handler Api Errors")
			}
		}
	}
}
