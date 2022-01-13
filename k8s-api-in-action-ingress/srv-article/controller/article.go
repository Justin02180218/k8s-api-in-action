package controller

import (
	"log"
	"net/http"
	"strconv"

	"com.justin.k8s.api/srv-article/model"
	"com.justin.k8s.api/srv-article/service"
	"github.com/gin-gonic/gin"
)

var articleService = service.NewArticleService()

func CreateArticle(ctx *gin.Context) {
	var article model.Article
	ctx.ShouldBindJSON(&article)

	var code int
	err := articleService.CreateArticle(&article)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": article,
		"msg":  "",
	})
}

func ArticleList(ctx *gin.Context) {
	userid, _ := strconv.Atoi(ctx.Param("userid"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))

	var code int
	articles, total, err := articleService.ArticleList(userid, pageSize, pageNum)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"data":  articles,
		"total": total,
		"msg":   "",
	})
}

func GetArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var code int
	article, err := articleService.GetArticle(id)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": article,
		"msg":  "",
	})
}

func EditArticle(ctx *gin.Context) {
	var article model.Article
	ctx.ShouldBindJSON(&article)

	id, _ := strconv.Atoi(ctx.Param("id"))

	var code int
	err := articleService.EditArticle(id, &article)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": article,
		"msg":  "",
	})
}

func DeleteArticle(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var code int
	err := articleService.DeleteArticle(id)
	if err != nil {
		log.Println(err)
		code = 500
	} else {
		code = 200
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": "",
		"msg":  "",
	})
}
