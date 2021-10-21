package handler

import (
	"github.com/labstack/echo/v4"
	"golang-simple-api/handler/request"
	"golang-simple-api/handler/response"
	"golang-simple-api/model"
	"golang-simple-api/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetArticle(c echo.Context) error {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	a, err := h.articleStore.GetById(id)
	if err != nil {
		return err
	}
	data := response.NewArticleResponse(a)
	return c.JSON(http.StatusOK, utils.NewResponse("Success get article", data))
}

func (h *Handler) Articles(c echo.Context) error {
	var (
		articles []model.Article
		_        int
	)
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 0
	}
	articles, _, err = h.articleStore.List(offset, limit)
	data := response.NewArticleListResponse(articles)
	return c.JSON(http.StatusOK, utils.NewResponse("Success get all article", data))
}

func (h *Handler) CreateArticle(c echo.Context) error {
	var a model.Article
	req := &request.ArticleCreateRequest{}
	if err := req.Bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := h.articleStore.CreateArticle(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	data := response.NewArticleResponse(&a)
	return c.JSON(http.StatusCreated, utils.NewResponse("Success create article", data))
}

func (h *Handler) UpdateArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	a, err := h.articleStore.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &request.ArticleUpdateRequest{}
	req.Populate(a)

	if err := req.Bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.articleStore.UpdateArticle(a); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	data := response.NewArticleResponse(a)
	return c.JSON(http.StatusOK, utils.NewResponse("Success update article", data))
}

func (h *Handler) DeleteArticle(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	a, err := h.articleStore.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	if err = h.articleStore.DeleteArticle(a); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	data := response.NewArticleResponse(a)
	return c.JSON(http.StatusOK, utils.NewResponse("Success update article", data))
}