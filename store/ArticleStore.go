package store

import (
	"errors"
	"golang-simple-api/model"
	"gorm.io/gorm"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (as *ArticleStore) GetById(id int) (*model.Article, error) {
	var m model.Article
	err := as.db.First(&m, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (as *ArticleStore) CreateArticle(a *model.Article) error  {
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ArticleStore) UpdateArticle(a *model.Article) error  {
	tx := as.db.Begin()
	if err := tx.Model(a).Updates(a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ArticleStore) DeleteArticle(a *model.Article) error  {
	return as.db.Delete(a).Error
}

func (as *ArticleStore) List(offset, limit int) ([]model.Article, int, error){
	var (
		articles []model.Article
		count int64
	)
	as.db.Model(&articles).Count(&count)
	as.db.Offset(offset).Limit(limit).Find(&articles)
	return articles, int(count), nil
}