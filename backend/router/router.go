package router

import (
	"cloud-note/controller/folder"
	"cloud-note/controller/user"
	"cloud-note/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)


func LoggerFunc(param gin.LogFormatterParams) string {

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
}
func StartServer() {
	gin.SetMode(gin.DebugMode)
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f,os.Stdout)

	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(LoggerFunc))
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	authGroup := r.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		authGroup.GET("hello", user.HelloAction)
		authGroup.GET("folder/all", folder.GetFoldersAction)
	}
	r.POST("/register", user.RegisterAction)
	r.POST("/signin", user.LoginAction)
	r.Run(":8001")
}
