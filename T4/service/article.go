package Service

import (
	"T4/model"
	"T4/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// @title 发布文章
// @description 获取文章内容传给下一级发布文章函数
// @param c *gin.Context "对应的上下文"
// @return code int "状态码"
func PostArticle(c *gin.Context) int {
	var msg model.Article
	err := c.ShouldBind(&msg)
	if err != nil {
		log.Println(err)
		return emsg.Error
	}
	code := model.PostArticle(&msg)
	return code
}

// @title 获取文章对象
// @description 获取文章id传给下一级获取文章函数
// @param c *gin.Context "对应的上下文"
// @return msg *model.Article "文章id对应的文章内容" code int "状态码"
func GetArticle(c *gin.Context) (*model.Article, int) {
	id, _ := strconv.Atoi(c.Param("id"))
	msg, code := model.GetArticle(id)
	return msg, code
}

// @title 删除文章
// @description 获取文章id传给下一级删除文章函数
// @param c *gin.Context "对应的上下文"
// @return code int "状态码"
func DeleteArticle(c *gin.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArticle(id)
	return code
}

// @title 点赞文章
// @description 获取文章id和当前用户名传给下一级点赞文章函数
// @param c *gin.Context "对应的上下文"
// @return code int "状态码"
func LikeArticle(c *gin.Context) int {
	username, err := c.Cookie("username")
	if err != nil {
		log.Println(err)
		return emsg.NoLogin
	}
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.LikeArticle(username, id)
	return code
}
