package controller

import (
	"ginFrame/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NoResponse 找不到路由 404
func NoResponse(c *gin.Context) {
	// 返回 404 状态码
	common.ServerError(c, http.StatusNotFound, "page not exists!")
}

// NoMethod 未知调用方式 405
func NoMethod(c *gin.Context) {
	common.ServerError(c, http.StatusMethodNotAllowed, "method not allowed")
}
