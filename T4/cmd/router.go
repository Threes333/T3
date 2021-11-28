package cmd

import (
	"T4/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	//用户相关
	r.POST("/register", controller.Register)
	r.GET("/login", controller.Login)
	r.GET("/home", controller.Home)
	//文章相关
	r.GET("/article/:id", controller.GetArticle)
	r.POST("/article", controller.PostArticle)
	r.DELETE("/article/:id", controller.DeleteArticle)
	r.POST("/article/likes/:id", controller.LikeArticle)
	_ = r.Run(":8080")
}
