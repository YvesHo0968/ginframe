package middleware

import (
	"fmt"
	"ginFrame/config"
	"github.com/gin-gonic/gin"
	"time"
)

// Logger 日志到文件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		config.Log.Info().
			Int("statusCode", statusCode).
			Str("latencyTime", fmt.Sprintf("%v", latencyTime)).
			Str("clientIP", clientIP).
			Str("reqMethod", reqMethod).
			Str("reqUri", reqUri).
			Str("proto", c.Request.Proto).
			Str("userAgent", c.Request.UserAgent()).
			Msg("request")
	}
}
