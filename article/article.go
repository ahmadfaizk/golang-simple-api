package article

import "golang-simple-api/model"

type Store interface {
	CreateArticle(article *model.Article) error
	UpdateArticle(article *model.Article) error
	DeleteArticle(article *model.Article) error
	List(offset, limit int) ([]model.Article, int, error)
	GetById(id int) (*model.Article, error)
}