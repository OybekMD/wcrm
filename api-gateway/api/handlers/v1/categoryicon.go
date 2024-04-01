package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "api-gateway/api/handlers/models"
	pbp "api-gateway/genproto/post"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"
)

// CategoryIcon
// @Summary     CreateCategoryIcon
// @Description Api for creating a new CategoryIcon
// @Tags        CategoryIcon
// @Accept      json
// @Produce     json
// @Param       CategoryIcon body models.CategoryIconReq true "CategoryIcon information for creating a new CategoryIcon"
// @Success     200 {object} models.CategoryIcon "Successfully created CategoryIcon"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/categoryIcon [post]
func (h *handlerV1) CreateCategoryIcon(c *gin.Context) {
	var (
		body        models.CategoryIcon
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

	response, err := h.serviceManager.PostService().CreateCategoryIcon(ctx, &pbp.CategoryIcon{
		Name:    body.Name,
		Picture: body.Picture,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create CategoryIcon", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary     ReadCategoryIcon
// @Description Api for getting a CategoryIcon by ID
// @Tags        CategoryIcon
// @Accept      json
// @Produce     json
// @Param       id query string true "CategoryIcon ID"
// @Success     200 {object} models.CategoryIcon "Successfully retrieved CategoryIcon"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/categoryIcon/:id [get]
func (h *handlerV1) ReadCategoryIcon(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ReadCategoryIcon(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get CategoryIcon", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary UpdateCategoryIcon
// @Description API for updating CategoryIcon by id
// @Tags CategoryIcon
// @Accept json
// @Produce json
// @Param body body models.CategoryIconUpdate true "CategoryIcon object to update"
// @Success 200 {object} models.CategoryIcon
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categoryIcon/:id [put]
func (h *handlerV1) UpdateCategoryIcon(c *gin.Context) {
	var (
		body        pbp.CategoryIcon
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

	response, err := h.serviceManager.PostService().UpdateCategoryIcon(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update CategoryIcon", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteCategoryIcon
// @Description API for deleting CategoryIcon by id
// @Tags CategoryIcon
// @Accept json
// @Produce json
// @Param id query string true "CategoryIcon ID"
// @Success 200 {object} models.CategoryIcon
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categoryIcon/:id [delete]
func (h *handlerV1) DeleteCategoryIcon(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeleteCategoryIcon(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete CategoryIcon", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListCategoryIcons returns list of CategoryIcons
// @Summary ListCategoryIcons
// @Description Api returns list of CategoryIcons
// @Tags CategoryIcon
// @Accept json
// @Produce json
// @Param page query int64 true "Page"
// @Param limit query int64 true "Limit"
// @Security ApiKeyAuth
// @Succes 200 {object} models.CategoryIcon
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categoryIcons/ [get]
func (h *handlerV1) ListCategoryIcons(c *gin.Context) {

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

	response, err := h.serviceManager.PostService().ListCategoryIcons(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list CategoryIcon", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
