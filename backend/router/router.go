package router

import (
	"cloud-note/controller"
	"cloud-note/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func StartServer() {
	gin.SetMode(gin.DebugMode)
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)

	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	authGroup := r.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		authGroup.GET("hello", controller.HelloAction)
	}
	r.POST("/register", controller.RegisterAction)
	r.POST("/signin", controller.LoginAction)
	r.Run(":8001")
}
