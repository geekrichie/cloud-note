package controller

import (
	"cloud-note/model"
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

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
	expirationTime := time.Now().Add(30 * time.Minute);

	claims :=&Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			//the expiry time is expressed
			ExpiresAt: expirationTime.Unix(),
		},
	}

   //Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
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
	username, err := c.Get("username")
	if err != true {
		log.Println("username is null")
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

func ParseToken(tokenString string) (string, error){
    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token)(interface{}, error) {
    	return jwtKey,nil
	})
    if err != nil {
    	return "", err
	}
	if !token.Valid {
		return "", err
	}
	return claims.Username,nil
}