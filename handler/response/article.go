package response

import (
	"golang-simple-api/model"
	"time"
)

type ArticleResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewArticleResponse(a *model.Article) *ArticleResponse {
	ar := new(ArticleResponse)
	ar.ID = a.ID
	ar.Title = a.Title
	ar.Body = a.Body
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt
	return ar
}

func NewArticleListResponse(articles []model.Article) []*ArticleResponse {
	arr := make([]*ArticleResponse, 0)
	for _, a := range articles {
		ar := new(ArticleResponse)
		ar.ID = a.ID
		ar.Title = a.Title
		ar.Body = a.Body
		ar.CreatedAt = a.CreatedAt
		ar.UpdatedAt = a.UpdatedAt
		arr = append(arr, ar)
	}
	return arr
}