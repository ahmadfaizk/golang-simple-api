package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(v1 *echo.Group)  {
	article := v1.Group("/article")
	article.GET("", h.Articles)
	article.POST("", h.CreateArticle)
	article.PUT("/:id", h.UpdateArticle)
	article.DELETE("/:id", h.DeleteArticle)
	article.GET("/:id", h.GetArticle)
}