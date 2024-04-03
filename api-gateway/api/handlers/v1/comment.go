package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "api-gateway/api/handlers/models"
	pbc "api-gateway/genproto/comment"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"
)

// Comment
// @Summary     CreateComment
// @Description Api for creating a new Comment
// @Security    BearerAuth
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Param       Comment body models.CommentReq true "Comment information for creating a new Comment"
// @Success     200 {object} models.Comment "Successfully created Comment"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/comment [post]
func (h *handlerV1) CreateComment(c *gin.Context) {
	var (
		body        models.CommentReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().CreateComment(ctx, &pbc.Comment{
		Content:   body.Content,
		UserId:    body.UserId,
		ProductId: body.ProductId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Comment", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary     ReadComment
// @Description Api for getting a Comment by ID
// @Security    BearerAuth
// @Tags        Comment
// @Accept      json
// @Produce     json
// @Param       id query string true "Comment ID"
// @Success     200 {object} models.Comment "Successfully retrieved Comment"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/comment/:id [get]
func (h *handlerV1) ReadComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().ReadComment(
		ctx, &pbc.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary UpdateComment
// @Description API for updating Comment by id
// @Security    BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param body body models.CommentUpdate true "Comment object to update"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment [put]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var (
		body        models.CommentUpdate
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().UpdateComment(ctx, &pbc.Comment{
		Id:      body.Id,
		Content: body.Content,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteComment
// @Description API for deleting Comment by id
// @Security    BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param id query string true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comment/:id [delete]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().DeleteComment(
		ctx, &pbc.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListComments returns list of Comments
// @Summary ListComments
// @Description Api returns list of Comments
// @Security    BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param page query int64 true "Page"
// @Param limit query int64 true "Limit"
// @Succes 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/ [get]
func (h *handlerV1) ListComments(c *gin.Context) {

	queryParams := c.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}

	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().ListComments(
		ctx, &pbc.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListCommentsByProductId returns list of Comments
// @Summary ListCommentsByProductId
// @Description Api returns list of Comments
// @Security    BearerAuth
// @Tags Comment
// @Accept json
// @Produce json
// @Param id query string true "Comment ID"
// @Succes 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/commentsbyproductid [get]
func (h *handlerV1) ListCommentsByProductId(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CommentService().ListCommentsByProductId(
		ctx, &pbc.IdRequest{
			Id: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Comment", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
