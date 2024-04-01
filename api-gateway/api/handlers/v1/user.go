package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"

	models "api-gateway/api/handlers/models"
	pbu "api-gateway/genproto/user"
	l "api-gateway/pkg/logger"
	"api-gateway/pkg/utils"
)

// @Summary		CreateUser
// @Description	Api for creating a new user
// @Tags		user
// @Accept		json
// @Produce		json
// @Param		User	body		models.User	true	"User"
// @Success		200		{object}	models.User
// @Failure		400		{object}	models.StandardErrorModel
// @Failure		500		{object}	models.StandardErrorModel
// @Router		/v1/user [post]
func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		body        models.User
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

	response, err := h.serviceManager.UserService().CreateUser(ctx, &pbu.User{
		Id:           body.Id,
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.Username,
		PhoneNumber:  body.PhoneNumber,
		Bio:          body.Bio,
		BirthDay:     body.BirthDay,
		Email:        body.Email,
		Avatar:       body.Avatar,
		Password:     body.Password,
		RefreshToken: body.RefreshToken,
		CreatedAt:    body.CreatedAt,
		UpdatedAt:    body.UpdatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}
	// response.RefreshToken = refresh
	c.JSON(http.StatusCreated, response)
}

// GetUser gets user by id
//
//	@Summary		GetUser
//	@Description	Api for getting user by id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"ID"
//	@Success		200	{object}	models.User
//	@Failure		400	{object}	models.StandardErrorModel
//	@Failure		500	{object}	models.StandardErrorModel
//	@Router			/v1/user/:id [get]
func (h *handlerV1) ReadUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().ReadUser(
		ctx, &pbu.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser updates a user by id
//
// @Summary UpdateUser
// @Description API for updating user by id
// @Tags user
// @Accept json
// @Produce json
// @Param body body models.User true "User object to update"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/:id [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
	var (
		body        pbu.User
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
	body.Id = c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteUser deletes a user by id
//
// @Summary DeleteUser
// @Description API for deleting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/user/:id [delete]
func (h *handlerV1) DeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().DeleteUser(
		ctx, &pbu.IdRequest{
			Id: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// ListUsers returns list of users
// @Summary ListUser
// @Description Api returns list of users
// @Tags user
// @Accept json
// @Produce json
// @Param page path int64 true "Page"
// @Param limit path int64 true "Limit"
// @Security ApiKeyAuth
// @Succes 200 {object} models.Users
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/ [get]
func (h *handlerV1) ListUsers(c *gin.Context) {

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

	response, err := h.serviceManager.UserService().ListUser(
		ctx, &pbu.GetAllRequest{
			Limit: params.Limit,
			Page:  params.Page,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
