package user

import (
	"cloud-note/model"
	"cloud-note/service/jwt"
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


func RegisterAction(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email    := c.PostForm("email")
	if username == "" || password == "" || email ==  ""{
		c.Writer.WriteHeader(http.StatusForbidden)
		return
	}
	w := md5.New()
	w.Write([]byte(password))
	md5password := fmt.Sprintf("%x", w.Sum(nil))
	user := &model.User{
		Username: username,
		Password: md5password,
		Email : email,
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


	tokenString, err := jwt.CreateToken(user.ID)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"username" : username,
			"token" : tokenString,
		},
	})

}


func HelloAction(c *gin.Context) {
	username, err := c.Get("uid")
	if err != true {
		log.Println("uid is null")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "hello",
		"data": gin.H{
			"msg" : "you jwt_token is valid",
			"username":username,
		},
	})
}
