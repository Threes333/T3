package controller

import (
	"T4/service"
	"T4/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

//发布文章的处理器函数
func PostArticle(c *gin.Context) {
	code := Service.PostArticle(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//获取文章的处理器函数
func GetArticle(c *gin.Context) {
	data, code := Service.GetArticle(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": *data,
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//删除文章的处理器函数
func DeleteArticle(c *gin.Context) {
	code := Service.DeleteArticle(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}

//点赞文章的处理器函数
func LikeArticle(c *gin.Context) {
	code := Service.LikeArticle(c)
	if code != emsg.Success {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  emsg.GetErrorMsg(code),
		})
	}
}
