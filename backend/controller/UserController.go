package controller

import (
	"cloud-note/model"
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterAction(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	w := md5.New()
	w.Write([]byte(password))
	md5password := fmt.Sprintf("%x", w.Sum(nil))
	user := &model.User{
		Username: username,
		Password: md5password,
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		InsertTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	status := user.Save()
	if status == true {
		c.JSON(http.StatusOK,gin.H{
			"code" : 1,
			"msg" : "创建用户成功",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1,
			"msg" : "创建用户失败",
		})
	}
}

func LoginAction(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
    user := model.User{}
    user.Find(username)
    if user.Username == "" {
		c.JSON(http.StatusOK,gin.H{
			"code" : -2,
			"msg" : "用户不存在",
		})
		return
	}
	w := md5.New()
	w.Write([]byte(password))
	md5password := fmt.Sprintf("%x", w.Sum(nil))
	if md5password != user.Password {
		c.JSON(http.StatusOK,gin.H{
			"code" : -1,
			"msg" : "密码错误",
		})
		return
	}
		c.JSON(http.StatusOK, gin.H{
			"code": -2,
			"msg":  "登录成功",
		})

}
