package request

import (
	"github.com/labstack/echo/v4"
	"golang-simple-api/model"
)

type ArticleCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Body string `json:"body" validate:"required"`
}

func (r *ArticleCreateRequest) Bind(c echo.Context, a *model.Article) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Title
	a.Body = r.Title
	return nil
}

type ArticleUpdateRequest struct {
	Title string `json:"title"`
	Body string `json:"body"`
}

func (r *ArticleUpdateRequest) Populate(a *model.Article) {
	r.Title = a.Title
	r.Body = a.Body
}

func (r *ArticleUpdateRequest) Bind(c echo.Context, a *model.Article) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Title
	a.Body = r.Body
	return nil
}