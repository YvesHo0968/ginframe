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

		errorToPrint := c.Errors.Last()
		if errorToPrint != nil {
			config.Log.Error().Interface("Errors", errorToPrint).Msg("Handler Errors")
			return
		}
	}
}
