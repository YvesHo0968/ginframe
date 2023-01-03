package ginFrame

import (
	"context"
	"fmt"
	"ginFrame/config"
	"ginFrame/route"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var g errgroup.Group
var GServer *Server
var Version string
var GoVersion string
var BuildTime string

type Server struct {
	GinServer *gin.Engine
}

func New() {
	// 初始化配置
	config.Init()

	if config.Viper.AppDebug {
		// 设置全局环境
		gin.SetMode(gin.DebugMode)
		gin.DefaultWriter = os.Stdout // 开启控制台打印
	} else {
		// 设置全局环境
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard // 关闭控制台打印
	}

	// 禁用控制台颜色
	gin.DisableConsoleColor()
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
			ReadTimeout:  60 * time.Second,
			WriteTimeout: 60 * time.Second,
		})

		gServer := servers[k]

		g.Go(func() error {
			err := gServer.ListenAndServe()

			if err != nil && err != http.ErrServerClosed {
				fmt.Println(err.Error())
				return err
			}

			return nil
		})

		fmt.Printf("Listen Server %s:%d\r\n", serverIp, v)
	}

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Println("Server Start Error", err)
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	// 等待中断信号，以优雅地关闭服务器
	quit := make(chan os.Signal)
	// 可以捕捉除了kill-9的所有中断信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	s := <-quit
	close(quit)

	// 接受到信号
	fmt.Println("Signal Received", s.String())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	for _, val := range servers {
		wg.Add(1)
		go func(val *http.Server) {
			defer wg.Done()

			if err := val.Shutdown(ctx); err != nil {
				fmt.Println(val.Addr, "Shutdown Error...")
			} else {
				fmt.Println(val.Addr, "Shutdown Success...")
			}
		}(val)
	}
	wg.Wait()

	fmt.Println("EXIT...")
}
