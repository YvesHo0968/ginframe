package controller

import (
	"github.com/gin-gonic/gin"
)

func NoResponse(c *gin.Context) {
	// 返回 404 状态码
	c.JSON(404, gin.H{
		"code": 404,
		"msg":  "page not exists!",
	})
}
