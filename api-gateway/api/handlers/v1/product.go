package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "api-gateway/api/handlers/models"
	pbp "api-gateway/genproto/post"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"
)

// Product
// @Summary     CreateProduct
// @Description Api for creating a new Product
// @Security    BearerAuth
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       Product body models.ProductReq true "Product information for creating a new Product"
// @Success     200 {object} models.Product "Successfully created Product"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/product [post]
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var (
		body        models.Product
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

	response, err := h.serviceManager.PostService().CreateProduct(ctx, &pbp.Product{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Product", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary     ReadProduct
// @Description Api for getting a Product by ID
// @Security    BearerAuth
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       id path string true "Product ID"
// @Success     200 {object} models.Product "Successfully retrieved Product"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/product/:id [get]
func (h *handlerV1) ReadProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ReadProduct(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get Product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary UpdateProduct
// @Description API for updating Product by id
// @Security    BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param	id path string true "Product ID"
// @Param body body models.ProductUpdate true "Product object to update"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/:id [put]
func (h *handlerV1) UpdateProduct(c *gin.Context) {
	var (
		body        pbp.Product
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
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
	iid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed while id to int64", l.Error(err))
		return
	}
	body.Id = int64(iid)

	response, err := h.serviceManager.PostService().UpdateProduct(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteProduct
// @Description API for deleting Product by id
// @Security    BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/product/:id [delete]
func (h *handlerV1) DeleteProduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeleteProduct(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProducts returns list of Products
// @Summary ListProducts
// @Description Api returns list of Products
// @Security    BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/products/ [get]
func (h *handlerV1) ListProducts(c *gin.Context) {

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

	response, err := h.serviceManager.PostService().ListProducts(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListProducts returns list of Products
// @Summary ListProducts
// @Description Api returns list of Products
// @Security    BearerAuth
// @Tags Product
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Succes 200 {object} models.Product
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/listproductwithcomments [get]
func (h *handlerV1) ListProductWithComment(c *gin.Context) {

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

	response, err := h.serviceManager.PostService().ListProductsWithComments(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Product", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
