package v1

import (
	"api-gateway/api/handlers/models"
	token "api-gateway/api/tokens"
	"api-gateway/config"
	pbu "api-gateway/genproto/user"
	"api-gateway/pkg/etc"
	l "api-gateway/pkg/logger"
	"context"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// Signin
// @Summary Signin
// @Description Signin - Api for Signining users
// @Tags Register
// @Accept json
// @Produce json
// @Param Signin body models.Signup true "Signup"
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/signup [post]
func (h *handlerV1) Signup(c *gin.Context) {
	var (
		body        models.Signup
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	err = body.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect email or password for validation")
		h.log.Error("Error validation", l.Error(err))
		return
	}

	response, err := h.serviceManager.UserService().CheckField(
		ctx, &pbu.CheckUser{
			Field: "email",
			Value: body.Email,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while getting exist fild",
		})
		h.log.Error(err.Error())
		return
	}
	if response.Exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email already use",
		})
		h.log.Error(err.Error())
		return
	}

	code := etc.GenerateCode(6)
	err = etc.SetRedis(code, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Generateing is wrong now. Please try again",
		})
	}
	etc.SendCode(body.Email, code)

	responsemessage := models.ResponseMessage{
		Message: "We send verification password to you email check your email!",
	}

	c.JSON(http.StatusOK, responsemessage)
}

// Login
// @Summary Login User
// @Description Login - Api for login users
// @Tags Register
// @Accept json
// @Produce json
// @Param login body models.LoginReq true "LoginReq"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login [post]
func (h *handlerV1) Login(c *gin.Context) {
	var (
		body        models.LoginReq
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	// Call your UserService.Login method with the retrieved email
	responseUser, err := h.serviceManager.UserService().Login(ctx, &pbu.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		h.log.Error(err.Error())
		return
	}

	// Handle login validation logic here...

	// Generate JWT tokens and respond
	// (Your existing code for generating tokens and responding)

	c.JSON(http.StatusOK, responseUser)
}

// Verification
// @Summary Verification User
// @Description LogIn - Api for verification users
// @Tags Register
// @Accept json
// @Produce json
// @Param verify body models.Verify true "Verify"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/verification [post]
func (h *handlerV1) Verification(c *gin.Context) {
	var (
		body        models.Verify
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	userdetail, err := etc.GetRedis(body.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Your code is expired",
		})
	}
	if body.Email != userdetail.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect email. Try again",
		})
		return
	}
	hashPassword, err := etc.HashPassword(userdetail.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Generateing Hash password is wrong now. Please try again",
		})
		h.log.Error("Error hashing password", l.Error(err))
		return
	}
	cfg := config.Load()

	// Create access and refresh tokens JWT
	h.jwthandler = token.JWTHandler{
		Sub:       userdetail.Email,
		Iss:       time.Now().String(),
		Exp:       cast.ToString(h.cfg.AccessTokenTimeout),
		Role:      "user",
		SigninKey: cfg.SignInKey,
		Timeout:   h.cfg.AccessTokenTimeout,
	}
	access, refresh, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "error generating token")
		return
	}

	createdUser, err := h.serviceManager.UserService().CreateUser(context.Background(), &pbu.User{
		Id:           uuid.New().String(),
		FirstName:    userdetail.FirstName,
		LastName:     userdetail.LastName,
		Email:        userdetail.Email,
		Password:     hashPassword,
		RefreshToken: refresh,
	})

	if createdUser == nil || err != nil {
		if err != nil {
			h.log.Error("Failed to create user", l.Error(err))
		} else {
			h.log.Error("Created user is nil")
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating user",
		})

		return
	}

	response := &models.UserResponse{
		Id:           uuid.New().String(),
		FirstName:    userdetail.FirstName,
		LastName:     userdetail.LastName,
		Email:        userdetail.Email,
		Password:     hashPassword,
		AccessToken:  access,
		RefreshToken: refresh,
		CreatedAt:    createdUser.CreatedAt,
		UpdatedAt:    createdUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

// ResetPassword
// @Summary ResetPassword User
// @Description ResetPassword - Api for ResetPassword users
// @Tags Register
// @Accept json
// @Produce json
// @Param resetpassword body models.ResetPassword true "ResetPassword"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/resetpassword [post]
func (h *handlerV1) ResetPassword(c *gin.Context) {
	var (
		body        models.ResetPassword
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// Check user isExist in db
	response, err := h.serviceManager.UserService().CheckField(
		ctx, &pbu.CheckUser{
			Field: "email",
			Value: body.Email,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while getting exist fild",
		})
		h.log.Error("Error while getting exist fild: ", l.Error(err))
		return
	}
	fmt.Println(body.Email)
	fmt.Println(response)

	if !response.Exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Email is not exist!",
		})
		return
	}

	uid := uuid.New().String()
	saveRedis := models.ResetChecker{
		Uid:   uid,
		Email: body.Email,
	}
	err = etc.SetRedisResetPassword(uid, saveRedis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Generateing is wrong now. Please try again",
		})
	}

	// Call your UserService.Login method with the retrieved email
	responseUser, err := h.serviceManager.UserService().GetFullName(ctx, &pbu.LoginRequest{
		Email: body.Email,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Call your UserService.Login method with the retrieved email"})
		h.log.Error(err.Error())
		return
	}

	etc.SendResetPasswordCode(body.Email, uid, responseUser.FirstName, responseUser.LastName)

	responsemessage := models.ResponseMessage{
		Message: "We sent code to you. Check your email!",
	}

	c.JSON(http.StatusOK, responsemessage)
}

// ResetCheckPassword
// @Summary ResetCheckPassword User
// @Description LogIn - Api for ResetCheckPassword users
// @Tags Register
// @Accept json
// @Produce json
// @Param resetchecker body models.ResetChecker true "ResetChecker"
// @Success 200 {object} models.ResponseMessage
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/resetcheckpassword [post]
func (h *handlerV1) ResetCheckPassword(c *gin.Context) {
	var (
		body        models.ResetChecker
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind json",
		})
		h.log.Error(err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	// err = body.Validate()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "Incorrect email for validation")
	// 	h.log.Error("Error Incorrect email for validation: ", l.Error(err))
	// 	return
	// }

	detail, err := etc.GetRedisResetPassword(body.Uid)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Your code is expired",
		})
		h.log.Error("Error hashing password", l.Error(err))
	}

	if detail.Uid != body.Uid {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Incorrect email. Try again",
		})
		return
	}
	hashPassword, err := etc.HashPassword(body.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Generateing Hash password is wrong now. Please try again",
		})
		h.log.Error("Error hashing password", l.Error(err))
		return
	}

	response, err := h.serviceManager.UserService().UpdatePassword(
		ctx, &pbu.UpdatePasswordRequest{
			Email:    body.Email,
			Password: hashPassword,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while changing password exist fild",
		})
		h.log.Error("Error while changing password exist fild: ", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}
