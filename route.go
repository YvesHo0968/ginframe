package ginFrame

import (
	"ginFrame/controller"
	"github.com/gin-gonic/gin"
)

func (s *Server) SetRoute() {
	r := s.ginServer

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

	// 设置不存在的路由
	r.NoRoute(controller.NoResponse)

	// 中间件
	//r.Use(middleWare.MyTime)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/test", controller.Test)
	}

	r.GET("/test", controller.Test)
	r.POST("/login", controller.Login)

	//r.POST("/somePost", posting)
	//r.PUT("/somePut", putting)
	//r.DELETE("/someDelete", deleting)
	//r.PATCH("/somePatch", patching)
	//r.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	return
}
