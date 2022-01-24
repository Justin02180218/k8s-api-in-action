package routers

import (
	"com.justin.k8s.api/srv-article/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	gin.SetMode(viper.GetString("server.type"))
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("article/create", controller.CreateArticle)
		v1.GET("article/list/:userid", controller.ArticleList)
		v1.GET("article/:id", controller.GetArticle)
		v1.PUT("article/:id", controller.EditArticle)
		v1.DELETE("article/:id", controller.DeleteArticle)
	}

	r.Run(":" + viper.GetString("server.port"))
}
