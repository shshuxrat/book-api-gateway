package handlers

import (
	"book-api-gateway/api/models"
	"book-api-gateway/genproto/book_service"
	"book-api-gateway/pkg/util"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBookCategory godoc
// @ID create-book-category
// @Router /v1/book_category [POST]
// @Summary create book category
// @Description Create Book Category
// @Tags book_category
// @Accept json
// @Produce json
// @Param book_category body models.CreateBookCategory true "book_category"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreateBookCategory(c *gin.Context) {
	var createBookCategory models.CreateBookCategory

	if err := c.BindJSON(&createBookCategory); err != nil {
		h.handleErrorResponse(c, 400, "wrong input or type ", err)
		return
	}

	resp, err := h.services.BookCategoryService().Create(
		context.Background(),
		&book_service.CreateBookCategory{
			Name: createBookCategory.Name,
		},
	)

	if !handleError(h.log, c, err, "error while creating book category") {
		return
	}

	h.handleSuccessResponse(c, 200, "created", resp)
}

// GetAllBookCategory godoc
// @ID get-all-book-category
// @Router /v1/book_category [GET]
// @Summary get all book category
// @Description Get All Book Category
// @Tags book_category
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllBookCategoryResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllBookCategory(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.BookCategoryService().GetAll(
		context.Background(),
		&book_service.GetAllBookCategoryRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Name:   c.Query("name"),
		},
	)
	if !handleError(h.log, c, err, "error  getting all attributes") {
		return
	}

	h.handleSuccessResponse(c, 200, "ok", resp)

}

// GetBookCategory godoc
// @ID get-book-category
// @Router /v1/book_category/{book_category_id} [GET]
// @Summary get book category by id
// @Description Get Book Category By Id
// @Tags book_category
// @Accept json
// @Produce json
// @Param book_category_id path string true "book_category_id"
// @Success 200 {object} models.ResponseModel{data=models.BookCategory} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetBookCategory(c *gin.Context) {
	var bookCategory models.BookCategory
	id := c.Param("book_category_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong uuid of book category", errors.New("wrong uuid of book category"))
		return
	}

	resp, err := h.services.BookCategoryService().GetById(
		context.Background(),
		&book_service.BookCategoryId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error getting attribute by id") {
		return
	}
	err = ParseToStruct(&bookCategory, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error parsing book category", err)
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", bookCategory)
}

// UpdateBookCategory godoc
// @ID update-book-category
// @Router /v1/book_category [PUT]
// @Summary update book category
// @Description Update Book Category
// @Tags book_category
// @Accept json
// @Produce json
// @Param book_category body models.UpdateBookCategory true "book_category"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdateBookCategory(c *gin.Context) {
	var bookCategory models.UpdateBookCategory
	if err := c.BindJSON(&bookCategory); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input book category", err)
		return
	}

	_, err := h.services.BookCategoryService().Update(
		context.Background(),
		&book_service.BookCategory{
			Id:   bookCategory.Id,
			Name: bookCategory.Name,
		},
	)
	if !handleError(h.log, c, err, "error while update book category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "updated", models.MsgModel{Msg: "Updated"})

}

// DeleteBookCategory godoc
// @ID delete-book-category
// @Router /v1/book_category/{book_category_id} [DELETE]
// @Summary delete book category by id
// @Description Delete Book Category By Id
// @Tags book_category
// @Accept json
// @Produce json
// @Param book_category_id path string true "book_category_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteBookCategory(c *gin.Context) {
	id := c.Param("book_category_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input uuid", errors.New("wrong input uuid"))
		return
	}
	_, err := h.services.BookCategoryService().Delete(
		context.Background(),
		&book_service.BookCategoryId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error while deleting book category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "deleted", models.MsgModel{Msg: "Deleted"})

}
