package middleware

import (
	"ginFrame/common"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Accept-Encoding, AccessToken, X-CSRF-Token, Authorization, Cache-Control, X-Requested-With, ResponseType, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

// RateLimiter token令牌限流
func RateLimiter() gin.HandlerFunc {
	//例如： 每秒产生500个令牌，最多存储200个令牌。
	l := rate.NewLimiter(500, 200)
	return func(c *gin.Context) {
		//当没有可用的令牌时返回false，也就是当没有可用的令牌时，禁止通行
		if !l.Allow() {
			common.ServerError(c, http.StatusTooManyRequests, "Rate Limit...")
			c.Abort()
		}
		//用可用的令牌时放行
		c.Next()
	}
}
