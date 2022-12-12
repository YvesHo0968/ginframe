package route

import (
	"ginFrame/common"
	"ginFrame/controller"
	"ginFrame/middleware"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func SetRoute(r *gin.Engine) {
	//r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	//
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		param.ClientIP,
	//		param.TimeStamp.Format(time.RFC1123),
	//		param.Method,
	//		param.Path,
	//		param.Request.Proto,
	//		param.StatusCode,
	//		param.Latency,
	//		param.Request.UserAgent(),
	//		param.ErrorMessage,
	//	)
	//}))
	//r.Use(gin.Recovery())

	//// 使用 Logger 中间件
	//r.Use(gin.Logger())
	//
	//// 使用 Recovery 中间件
	//r.Use(gin.Recovery())
	r.Use(favicon.New("./favicon.ico"))
	//r.StaticFile("/favicon.ico", "./favicon.ico")

	// 设置不存在的路由
	r.NoRoute(controller.NoResponse)

	// 设置位置调用方式
	r.NoMethod(controller.NoMethod)

	// 中间件
	//r.Use(middleware.MyTime)

	// 跨域
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"22":      common.FilePath(),
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/test", controller.Test)
	}

	r.GET("/test", controller.Test)
	r.POST("/login", controller.Login)

	// 加载资源文件
	r.Static("/static", "./static")

	// 路由重定向
	r.GET("/redirect", func(context *gin.Context) {
		// 重定向 301 http.StatusMovedPermanently
		context.Redirect(http.StatusMovedPermanently, "https://www.bilibili.com")
	})

	//r.POST("/somePost", posting)
	//r.PUT("/somePut", putting)
	//r.DELETE("/someDelete", deleting)
	//r.PATCH("/somePatch", patching)
	//r.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	return
}
