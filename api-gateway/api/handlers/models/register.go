package models

import (
	// "api-gateway/api/handlers/models"
	pbu "api-gateway/genproto/user"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// LoginReq ...
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserResponse ...
type UserResponse struct {
	Id           string `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"udpated_at"`
}

// Signup ...
type Signup struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// ResponseMessage ...
type ResponseMessage struct {
	Message string `json:"content"`
}

func (rm *Signup) Validate() error {
	return validation.ValidateStruct(
		rm,
		validation.Field(&rm.Email, validation.Required, is.Email),
		validation.Field(
			&rm.Password,
			validation.Required,
			validation.Length(8, 30),
			validation.Match(regexp.MustCompile("[a-z]|[A-Z][1-9]")),
		),
	)
}

func ParseStruct(req *pbu.User, access string) *User {
	return &User{
		Id:           req.Id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Username:     req.Username,
		Bio:          req.Bio,
		BirthDay:     req.BirthDay,
		Email:        req.Email,
		Password:     req.Password,
		Avatar:       req.Avatar,
		AccessToken:  access,
		RefreshToken: req.RefreshToken,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}
}

type ResetPassword struct {
	Email string `json:"email"`
}

type Verify struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type ResetChecker struct {
	Uid         string `json:"uid"`
	Email       string `json:"email"`
	NewPassword string `json:"newpassword"`
}
