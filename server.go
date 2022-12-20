package ginFrame

import (
	"fmt"
	"ginFrame/config"
	"ginFrame/route"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

var (
	g       errgroup.Group
	GServer *Server
)

type Server struct {
	GinServer *gin.Engine
}

func New() {
	// 初始化配置
	config.Init()

	if config.Viper.AppDebug {
		// 设置全局环境
		gin.SetMode(gin.DebugMode)
	} else {
		// 设置全局环境
		gin.SetMode(gin.ReleaseMode)
	}

	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//f, _ := os.OpenFile("./gin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	GServer = &Server{
		GinServer: gin.Default(),
	}

	// 设置路由
	route.SetRoute(GServer.GinServer)

	//ee := config.Rdb.Set(config.RdbCtx, "qwqww12121", "value", time.Minute*2).Err()
	//
	//if ee != nil {
	//	fmt.Println("错误")
	//}

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

	// 单端口启动
	//GServer.GinServer.Run(":8080") // listen and serve on 0.0.0.0:8080

	//GServer.GinServer.RunListener()

	ports := config.Viper.Ports       // 监听端口
	serverIp := config.Viper.ServerIp // 监听ip

	var servers []*http.Server

	for k, v := range ports {
		servers = append(servers, &http.Server{
			Addr:         fmt.Sprintf("%s:%d", serverIp, v),
			Handler:      GServer.GinServer,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		})

		gServer := servers[k]

		g.Go(func() error {
			err := gServer.ListenAndServe()
			if err != nil && err != http.ErrServerClosed {
				log.Print(err)
			}
			return err
		})

		fmt.Printf("Listen port %s:%d\r\n", serverIp, v)
	}

	if err := g.Wait(); err != nil {
		log.Print(err)
	}

}
