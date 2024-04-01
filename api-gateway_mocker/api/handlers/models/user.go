package models

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	Id              string `json:"id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Username        string `json:"username"`
	PhoneNumber     string `json:"phone_number"`
	Bio             string `json:"bio"`
	BirthDay        string `json:"birth_day"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	ExperienceLevel int64  `json:"experiience_level"`
	Coint           int64  `json:"coint"`
	Score           int64  `json:"score"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updeted_at"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(5, 15), validation.Match(regexp.MustCompile("[a-z]|[A-Z][0-9]"))),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 50), validation.Match(regexp.MustCompile("^[A-Z][a-z]*$"))),
	)
}
