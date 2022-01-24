package service

import (
	"com.justin.k8s.api/srv-article/dao"
	"com.justin.k8s.api/srv-article/model"
)

type ArticleService interface {
	CreateArticle(article *model.Article) error
	ArticleList(userid int, pageSize, pageNum int) ([]model.Article, int, error)
	GetArticle(id int) (model.Article, error)
	EditArticle(id int, article *model.Article) error
	DeleteArticle(id int) error
}

type ArticleServiceImpl struct {
	ArticleDao dao.ArticleDao
}

func NewArticleService() ArticleService {
	return &ArticleServiceImpl{
		ArticleDao: dao.NewArticleDao(),
	}
}

func (a *ArticleServiceImpl) CreateArticle(article *model.Article) error {
	return a.ArticleDao.CreateArticle(article)
}

func (a *ArticleServiceImpl) ArticleList(userid int, pageSize, pageNum int) ([]model.Article, int, error) {
	return a.ArticleDao.GetArticlesByUserID(userid, pageSize, pageNum)
}

func (a *ArticleServiceImpl) GetArticle(id int) (model.Article, error) {
	return a.ArticleDao.GetArticle(id)
}

func (a *ArticleServiceImpl) EditArticle(id int, article *model.Article) error {
	return a.ArticleDao.UpdateArticle(id, article)
}

func (a *ArticleServiceImpl) DeleteArticle(id int) error {
	return a.ArticleDao.DeleteArticle(id)
}
