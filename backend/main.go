package main

import (
	"cloud-note/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", controller.RegisterAction)
	r.POST("/signin", controller.LoginAction)
	r.Run(":8001")
}
