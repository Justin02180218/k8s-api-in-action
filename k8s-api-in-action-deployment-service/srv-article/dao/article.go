package dao

import (
	"com.justin.k8s.api/common/databases"
	"com.justin.k8s.api/srv-article/model"
)

type ArticleDao interface {
	CreateArticle(article *model.Article) error
	GetArticlesByUserID(userID int, pageSize, pageNum int) ([]model.Article, int, error)
	GetArticle(id int) (model.Article, error)
	UpdateArticle(id int, article *model.Article) error
	DeleteArticle(id int) error
}

type ArticleDaoImpl struct{}

func NewArticleDao() ArticleDao {
	return &ArticleDaoImpl{}
}

func (a *ArticleDaoImpl) CreateArticle(article *model.Article) error {
	return databases.DB.Create(article).Error
}

func (a *ArticleDaoImpl) GetArticlesByUserID(userID int, pageSize, pageNum int) ([]model.Article, int, error) {
	var articles []model.Article
	var total int

	err := databases.DB.Limit(pageSize).Offset((pageNum-1)*pageSize).Where("userid = ?", userID).Find(&articles).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return articles, total, nil
}

func (a *ArticleDaoImpl) GetArticle(id int) (model.Article, error) {
	var article model.Article
	err := databases.DB.Where("id = ?", id).First(&article).Error
	return article, err
}

func (a *ArticleDaoImpl) UpdateArticle(id int, article *model.Article) error {
	return databases.DB.Model(&model.Article{}).Where("id = ?", id).Update(article).Error
}

func (a *ArticleDaoImpl) DeleteArticle(id int) error {
	return databases.DB.Where("id = ?", id).Delete(&model.Article{}).Error
}
