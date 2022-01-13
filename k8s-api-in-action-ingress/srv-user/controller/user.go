package controller

import (
	"log"
	"net/http"

	"com.justin.k8s.api/srv-user/model"
	"com.justin.k8s.api/srv-user/service"
	"github.com/gin-gonic/gin"
)

var userService = service.NewUserService()

func Register(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	var code int
	err := userService.Register(&user)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": user,
		"msg":  "",
	})
}

func Login(ctx *gin.Context) {
	var user model.User
	ctx.ShouldBindJSON(&user)

	var code int
	u, err := userService.Login(user.Phone, user.Password)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": u,
		"msg":  "",
	})
}
