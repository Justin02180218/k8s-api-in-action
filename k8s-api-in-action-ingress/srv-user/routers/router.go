package routers

import (
	"com.justin.k8s.api/srv-user/controller"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	gin.SetMode(viper.GetString("server.type"))
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("regist", controller.Register)
		v1.POST("login", controller.Login)
	}

	r.Run(":" + viper.GetString("server.port"))
}
