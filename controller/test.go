package controller

import (
	"ginFrame/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	name := c.DefaultQuery("name", "122")
	action := c.Query("action")
	message := name + " is " + action
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    action,
		"ip":      c.ClientIP(),
	})
	//c.String(http.StatusOK, message)
}

// LoginData 绑定为json
type LoginData struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var json LoginData
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if json.User != "manu" || json.Password != "123" {
		//c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		//return
	}

	// You also can use a struct
	var msg struct {
		Name    string `json:"user"`
		Message string `json:"message"`
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// Note that msg.Name becomes "user" in the JSON
	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	//ginFrame.Success(c, msg)
	common.Success(c, msg)
	//c.JSON(http.StatusOK, msg)

	//c.JSON(http.StatusOK, gin.H{
	//	"user":     json.User,
	//	"password": json.Password,
	//})
}
