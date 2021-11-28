package Service

import (
	"T4/model"
	"T4/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @title 用户注册
// @description 获取用户名和密码传给下一级用户注册函数
// @param c *gin.Context "相应的上下文"
// @return code int "状态码"
func Register(c *gin.Context) int {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	code := model.Register(user.UserName, user.PassWord)
	return code
}

// @title 用户登录
// @description 获取用户名和密码传给下一级用户登录函数
// @param c *gin.Context "相应的上下文"
// @return code int "状态码"
func Login(c *gin.Context) int {
	username := c.Query("username")
	password := c.Query("password")
	code := model.Login(username, password)
	if code == emsg.Success {
		cookie := &http.Cookie{
			Name:     "username",
			Value:    username,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		//c.SetCookie("username", username, 3600, "/", "127.0.0.1", false, true)
	}
	return code
}
