package ginFrame

import (
	"ginFrame/config"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Server struct {
	ginServer *gin.Engine
}

func New() {
	config.InitRedis()
	gin.SetMode(gin.DebugMode)

	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	s := &Server{
		ginServer: gin.Default(),
	}

	// 设置路由
	s.SetRoute()

	//fmt.Println(dd)

	//err := dd.Set("key1", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := config.Redis.Get("key1").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)

	s.ginServer.Run(":8080") // listen and serve on 0.0.0.0:8080

	return
}
