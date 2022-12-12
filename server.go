package ginFrame

import (
	"fmt"
	"ginFrame/config"
	"ginFrame/route"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

var GServer *Server

type Server struct {
	GinServer *gin.Engine
}

func New() {
	// 启动redis
	config.InitRedis()

	// 启动数据库
	config.InitDb()

	// 设置全局环境
	gin.SetMode(gin.DebugMode)

	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	f, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	GServer = &Server{
		GinServer: gin.Default(),
	}

	// 设置路由
	route.SetRoute(GServer.GinServer)

	ee := config.Rdb.Set("qwqww12121", "value", 0).Err()

	if ee != nil {
		fmt.Println("错误")
	}

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

	GServer.GinServer.Run(":8080") // listen and serve on 0.0.0.0:8080
	//GServer.GinServer.RunListener()

	//server01 := &http.Server{
	//	Addr:         ":8080",
	//	Handler:      GServer.GinServer,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//
	//server02 := &http.Server{
	//	Addr:         ":8081",
	//	Handler:      GServer.GinServer,
	//	ReadTimeout:  5 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//}
	//
	//server01.ListenAndServe()
	//server02.ListenAndServe()

	return
}
