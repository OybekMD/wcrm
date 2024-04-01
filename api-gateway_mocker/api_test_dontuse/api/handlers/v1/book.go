package v1

import (
	"api-gateway/api/model"
	pb "api-gateway/genproto/book"
	"api-gateway/pkg/logger"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreateBook
// @Summary create Book
// @Tags Book
// @Description Insert a new Book with provided details
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param BookDetails body model.Item true "Create Book"
// @Success 201 {object} model.Item
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/create [post]
func (h *handlerV1) CreateBook(c *gin.Context) {
	var (
		body       model.Item
		jspMarshal protojson.MarshalOptions
	)
	jspMarshal.UseProtoNames = true

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	if body.Amount < 0 {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: "amount cannot be smaller than zero",
		})
		return
	}

	if body.Price < 0 {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: "price cannot be smaller than zero",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	id := uuid.New().String()
	resp, err := h.serviceManager.MockBookService().CreateBook(ctx, &pb.Book{
		Id:          id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Amount:      int64(body.Amount),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.log.Error("error while creating Book", logger.Error(err))
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update Book
// @Summary update Book
// @Tags Book
// @Description Update ptoduct
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param UserInfo body model.Item true "Update Book"
// @Success 201 {object} model.Item
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/update/{id} [put]
func (h *handlerV1) UpdateBook(c *gin.Context) {
	var (
		body        pb.Book
		jspbMarshal protojson.MarshalOptions
	)
	id := c.Param("id")

	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: err.Error(),
		})
		h.log.Error("cannot bind json", logger.Error(err))
		return
	}

	if body.Amount < 0 {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: "0 dan kichik amount kiritib bo`lmaydi",
		})
		return
	}

	if body.Price < 0 {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: "price cannot be smaller than zero",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MockBookService().UpdateBook(ctx, &pb.Book{
		Id:          id,
		Name:        body.Name,
		Description: body.Description,
		Price:       body.Price,
		Amount:      body.Amount,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.log.Error("error while updating Book", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get Book By Id
// @Summary get Book by id
// @Tags Book
// @Description Get Book
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Id"
// @Success 201 {object} model.Item
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/get/{id} [get]
func (h *handlerV1) GetBookById(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MockBookService().GetBookById(ctx, &pb.BookId{
		BookId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})
		h.log.Error("cannot get Book", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Delete Book
// @Summary delete Book
// @Tags Book
// @Description Delete Book
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 201 {object} model.Status
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/delete/{id} [delete]
func (h *handlerV1) DeleteBook(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MockBookService().DeleteBook(ctx, &pb.BookId{
		BookId: id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})

		h.log.Error("cannot delete Book", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// Get All Books
// @Summary get all Books
// @Tags Book
// @Description get all Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param page path string true "page"
// @Param limit path string true "limit"
// @Success 201 {object} model.ListItems
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/{page}/{limit} [get]
func (h *handlerV1) ListBooks(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	page := c.Param("page")
	intpage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: err.Error(),
		})
		h.log.Error("cannot parse page query param", logger.Error(err))
		return
	}

	limit := c.Param("limit")
	intlimit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{
			Code:  ErrorBadRequest,
			Error: err.Error(),
		})
		h.log.Error("cannot parse limit query param", logger.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.MockBookService().ListBooks(ctx, &pb.GetAllBookRequest{
		Page:  int64(intpage),
		Limit: int64(intlimit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ResponseError{
			Code:  ErrorCodeInternalServerError,
			Error: err.Error(),
		})

		h.log.Error("cannot list Books", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// // Get All Purchased Books by user
// // @Summary get all purchased Books by user id
// // @Tags Book
// // @Description get all purchased Books by user id
// // @Security BearerAuth
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Success 201 {object} model.BoughtItemsList
// // @Failure 400 string Error models.ResponseError
// // @Failure 500 string Error models.ResponseError
// // @Router /v1/books/get/{id} [get]
// func (h *handlerV1) GetPurchasedBooksByUserId(c *gin.Context) {
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	userId := c.Param("id")
// 	if userId == "" {
// 		fmt.Println("=======")
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	response, err := h.serviceManager.MockBookService().GetBoughtBooksByUserId(ctx, &pb.UserId{
// 		UserId: userId,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ResponseError{
// 			Code:  ErrorCodeInternalServerError,
// 			Error: err.Error(),
// 		})

// 		h.log.Error("cannot list Books purchased by user", logger.Error(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// Buy Book
// @Summary buy a Book
// @Tags Book
// @Description buy a Book
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param PurchaseInfo body model.BuyItemRequest true "Purchase a Book"
// @Success 201 {object} model.BuyItemResponse
// @Failure 400 string Error models.ResponseError
// @Failure 500 string Error models.ResponseError
// @Router /v1/book/buy [post]
// func (h *handlerV1) BuyBook(c *gin.Context) {
// 	var (
// 		res         model.BuyItemResponse
// 		body        model.BuyItemRequest
// 		jspbMarshal protojson.MarshalOptions
// 	)

// 	jspbMarshal.UseProtoNames = true
// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, model.ResponseError{
// 			Code:  ErrorCodeInvalidJSON,
// 			Error: err.Error(),
// 		})
// 		h.log.Error("cannot bind json", logger.Error(err))
// 		return
// 	}

// 	if body.Amount < 0 {
// 		c.JSON(http.StatusBadRequest, model.ResponseError{
// 			Code:  ErrorBadRequest,
// 			Error: "0 dan kichik amount kiritib bo`lmaydi",
// 		})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	status, err := h.serviceManager.MockBookService().CheckAmount(ctx, &pb.BookId{BookId: body.BookId})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ResponseError{
// 			Code:  ErrorCodeInternalServerError,
// 			Error: err.Error(),
// 		})

// 		h.log.Error("cannot list Books purchased by user", logger.Error(err))
// 		return
// 	}
// 	if status.Amount == 0 {
// 		c.JSON(http.StatusBadRequest, model.ResponseError{
// 			Code:  ErrorBadRequest,
// 			Error: "the Book is not currently available, sorry",
// 		})

// 		h.log.Error("not available Book", logger.Error(err))
// 		return

// 	}
// 	if !(status.Amount < body.Amount) {
// 		c.JSON(http.StatusBadRequest, model.ResponseError{
// 			Code:  ErrorCodeInternalServerError,
// 			Error: fmt.Sprintf("we have only %d, sorry", status.Amount),
// 		})

// 		h.log.Error("not enough", logger.Error(err))
// 		return
// 	}
// 	buyResp, err := h.serviceManager.MockBookService().BuyBook(ctx, &pb.BuyBookRequest{
// 		UserId:    body.UserId,
// 		BookId: body.BookId,
// 		Amount:    body.Amount,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ResponseError{
// 			Code:  ErrorCodeInternalServerError,
// 			Error: err.Error(),
// 		})

// 		h.log.Error("cannot purchase the Book", logger.Error(err))
// 		return
// 	}
// 	_, err = h.serviceManager.MockBookService().DecreaseAmount(ctx, &pb.BookAmountRequest{
// 		BookId: body.BookId,
// 		Amount:    body.Amount,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, model.ResponseError{
// 			Code:  ErrorCodeInternalServerError,
// 			Error: err.Error(),
// 		})
// 		h.log.Error("cannot decrease the amount", logger.Error(err))
// 		return
// 	}

// 	res.Message = "successfully purchased"
// 	res.BookId = body.BookId
// 	res.UserId = body.UserId
// 	res.Amount = body.Amount
// 	res.BookName = buyResp.Name

// 	c.JSON(http.StatusOK, res)
// }
