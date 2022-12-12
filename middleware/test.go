package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// MyTime 定义中间
func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	// 统计时间
	since := time.Since(start)
	fmt.Println("程序用时：", since)
}
