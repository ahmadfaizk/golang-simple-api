package handler

import "golang-simple-api/article"

type Handler struct {
	articleStore article.Store
}

func NewHandler(as article.Store) *Handler {
	return &Handler{
		articleStore: as,
	}
}