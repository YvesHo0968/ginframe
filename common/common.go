package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	type SuccessData struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}

	d := SuccessData{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}

	c.JSON(http.StatusOK, d)
}

func Uuid() string {
	return uuid.New().String()
}
