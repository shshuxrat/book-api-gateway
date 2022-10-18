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

// CreateBook godoc
// @ID create-book
// @Router /v1/book [POST]
// @Summary create book
// @Description Create Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.CreateBook true "book"
// @Success 200 {object} models.ResponseModel{data=string} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) CreateBook(c *gin.Context) {
	var book models.CreateBook
	if err := c.BindJSON(&book); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input for book", err)
		return
	}
	resp, err := h.services.BookService().Create(
		context.Background(),
		&book_service.CreateBook{
			Name:       book.Name,
			CategoryId: book.Category_id,
		},
	)

	if !handleError(h.log, c, err, "error while creating book") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "created", resp)
}

// GetAllBook godoc
// @ID get-all-book
// @Router /v1/book [GET]
// @Summary get all book
// @Description Get All Book
// @Tags book
// @Accept json
// @Produce json
// @Param name query string false "name"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} models.ResponseModel{data=models.GetAllBookResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetAllBook(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "100")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}
	resp, err := h.services.BookService().GetAll(
		context.Background(),
		&book_service.GetAllBookRequest{
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

// GetBook godoc
// @ID get-book
// @Router /v1/book/{book_id} [GET]
// @Summary get book by id
// @Description Get Book By Id
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} models.ResponseModel{data=models.GetBookResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) GetBook(c *gin.Context) {
	var bookData models.GetBookResponse
	id := c.Param("book_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input  book id", errors.New("wrong input  book id"))
		return
	}

	resp, err := h.services.BookService().GetById(
		context.Background(),
		&book_service.BookId{
			Id: id,
		},
	)

	if !handleError(h.log, c, err, "error while getting book") {
		return
	}
	err = ParseToStruct(&bookData, resp)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error parsing bookdata", err)
	}
	h.handleSuccessResponse(c, http.StatusOK, "ok", bookData)
}

// UpdateBook godoc
// @ID update-book
// @Router /v1/book [PUT]
// @Summary update book
// @Description Update Book
// @Tags book
// @Accept json
// @Produce json
// @Param book body models.UpdateBook true "book"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) UpdateBook(c *gin.Context) {
	var updateBook models.UpdateBook
	if err := c.BindJSON(&updateBook); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "wrong input for update book", err)
		return
	}
	_, err := h.services.BookService().Update(
		context.Background(),
		&book_service.UpdateBook{
			Id:         updateBook.Id,
			CategoryId: updateBook.Category_id,
			Name:       updateBook.Name,
		},
	)

	if !handleError(h.log, c, err, "error while updating book") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "updated", models.MsgModel{Msg: "Updated"})
}

// DeleteBook godoc
// @ID delete-book
// @Router /v1/book/{book_id} [DELETE]
// @Summary delete book by id
// @Description Delete Book By Id
// @Tags book
// @Accept json
// @Produce json
// @Param book_id path string true "book_id"
// @Success 200 {object} models.ResponseModel{data=models.MsgModel} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handler) DeleteBook(c *gin.Context) {
	id := c.Param("book_id")
	if !util.IsValidUUID(id) {
		h.handleErrorResponse(c, http.StatusBadGateway, "wrong input book id", errors.New("wrong input book id"))
		return
	}
	_, err := h.services.BookService().Delete(
		context.Background(),
		&book_service.BookId{
			Id: id,
		},
	)
	if !handleError(h.log, c, err, "error while deleting book") {
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "deleted", models.MsgModel{Msg: "Deleted"})
}
