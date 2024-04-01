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

// Orderproduct
// @Summary     CreateOrderproduct
// @Description Api for creating a new Orderproduct
// @Tags        Orderproduct
// @Accept      json
// @Produce     json
// @Param       Orderproduct body models.OrderproductReq true "Orderproduct information for creating a new Orderproduct"
// @Success     200 {object} models.Orderproduct "Successfully created Orderproduct"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/orderproduct [post]
func (h *handlerV1) CreateOrderproduct(c *gin.Context) {
	var (
		body        models.Orderproduct
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

	response, err := h.serviceManager.PostService().CreateOrderproduct(ctx, &pbp.Orderproduct{
		UserId:    body.UserId,
		ProductId: body.ProductId,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create Orderproduct", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// @Summary     ReadOrderproduct
// @Description Api for getting a Orderproduct by ID
// @Tags        Orderproduct
// @Accept      json
// @Produce     json
// @Param       id query string true "Orderproduct ID"
// @Success     200 {object} models.Orderproduct "Successfully retrieved Orderproduct"
// @Failure     400 {object} models.StandardErrorModel "Bad request"
// @Failure     500 {object} models.StandardErrorModel "Internal server error"
// @Router      /v1/orderproduct/:id [get]
func (h *handlerV1) ReadOrderproduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().ReadOrderproduct(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get Orderproduct", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary UpdateOrderproduct
// @Description API for updating Orderproduct by id
// @Tags Orderproduct
// @Accept json
// @Produce json
// @Param	id query string true "Orderproduct ID"
// @Param body body models.OrderproductUpdate true "Orderproduct object to update"
// @Success 200 {object} models.Orderproduct
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orderproduct/:id [put]
func (h *handlerV1) UpdateOrderproduct(c *gin.Context) {
	var (
		body        pbp.Orderproduct
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")
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

	response, err := h.serviceManager.PostService().UpdateOrderproduct(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update Orderproduct", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// @Summary DeleteOrderproduct
// @Description API for deleting Orderproduct by id
// @Tags Orderproduct
// @Accept json
// @Produce json
// @Param id query string true "Orderproduct ID"
// @Success 200 {object} models.Orderproduct
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orderproduct/:id [delete]
func (h *handlerV1) DeleteOrderproduct(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeleteOrderproduct(
		ctx, &pbp.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete Orderproduct", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListOrderproducts returns list of Orderproducts
// @Summary ListOrderproducts
// @Description Api returns list of Orderproducts
// @Tags Orderproduct
// @Accept json
// @Produce json
// @Param page query int64 true "Page"
// @Param limit query int64 true "Limit"
// @Security ApiKeyAuth
// @Succes 200 {object} models.Orderproduct
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/orderproducts/ [get]
func (h *handlerV1) ListOrderproducts(c *gin.Context) {

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

	response, err := h.serviceManager.PostService().ListOrderproducts(
		ctx, &pbp.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list Orderproduct", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
