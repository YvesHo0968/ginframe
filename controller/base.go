package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoResponse(c *gin.Context) {
	// 返回 404 状态码
	c.JSON(404, gin.H{
		"code": http.StatusNotFound,
		"msg":  "page not exists!",
	})
}

// NoMethod 未知调用方式
func NoMethod(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{
		"code": http.StatusMethodNotAllowed,
		"msg":  "method not allowed",
	})
}
