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

// Category
// @Summary     CreateCategory
// @Description Api for creating a new Category
// @Security    BearerAuth
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       Category body models.CategoryReq true "Category information for creating a new Category"
// @Success     200 {object} models.Category "Successfully created Category"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/category [post]
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var (
		body        models.Category
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

	response, err := h.serviceManager.PostService().CreateCategory(ctx, &pbp.Category{
		Name:   body.Name,
		IconId: body.IconId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Category", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary     ReadCategory
// @Description Api for getting a Category by ID
// @Security    BearerAuth
// @Tags        Category
// @Accept      json
// @Produce     json
// @Param       id query string true "Category ID"
// @Success     200 {object} models.Category "Successfully retrieved Category"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/category/:id [get]
func (h *handlerV1) ReadCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ReadCategory(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get Category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary UpdateCategory
// @Description API for updating Category by id
// @Security    BearerAuth
// @Tags Category
// @Accept json
// @Produce json
// @Param body body models.CategoryUpdate true "Category object to update"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/category/:id [put]
func (h *handlerV1) UpdateCategory(c *gin.Context) {
	var (
		body        pbp.Category
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

	response, err := h.serviceManager.PostService().UpdateCategory(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteCategory
// @Description API for deleting Category by id
// @Security    BearerAuth
// @Tags Category
// @Accept json
// @Produce json
// @Param id query string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/category/:id [delete]
func (h *handlerV1) DeleteCategory(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeleteCategory(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListCategorys returns list of Categorys
// @Summary ListCategorys
// @Description Api returns list of Categorys
// @Security    BearerAuth
// @Tags Category
// @Accept json
// @Produce json
// @Param page query int64 true "Page"
// @Param limit query int64 true "Limit"
// @Succes 200 {object} models.Category
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/categorys/ [get]
func (h *handlerV1) ListCategorys(c *gin.Context) {

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

	response, err := h.serviceManager.PostService().ListCategorys(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Category", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
