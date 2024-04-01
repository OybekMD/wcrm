package postgres

import (
	"testing"
	"user-service/config"
	pbu "user-service/genproto/user"
	"user-service/pkg/db"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserMockPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	// Test Create User
	repo := NewUserRepoMock(db)
	id := uuid.New().String()
	user := &pbu.User{
		Id:           id,
		FirstName:    "Oybek",
		LastName:     "Atamatov",
		Username:     "oybekmd",
		PhoneNumber:  "998999790445",
		Bio:          "Test bio",
		BirthDay:     "2003-08-01",
		Email:        "testemail@gamil.com",
		Avatar:       "example.com/user.jpg",
		Password:     "hello1234",
		RefreshToken: "testToken",
	}
	createdUser, err := repo.CreateUserDB(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, createdUser.Id)
	assert.Equal(t, user.FirstName, createdUser.FirstName)
	assert.Equal(t, user.LastName, createdUser.LastName)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Bio, createdUser.Bio)
	assert.Equal(t, user.BirthDay, createdUser.BirthDay)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.Avatar, createdUser.Avatar)
	assert.NotNil(t, createdUser.CreatedAt)
	assert.NotNil(t, createdUser.UpdatedAt)

	// Test Read User
	getUser, err := repo.ReadUserDB(&pbu.IdRequest{Id: user.Id})
	assert.NoError(t, err)
	assert.Equal(t, user.Id, getUser.Id)
	assert.Equal(t, user.FirstName, getUser.FirstName)
	assert.Equal(t, user.LastName, getUser.LastName)
	assert.Equal(t, user.Username, getUser.Username)
	assert.Equal(t, user.Bio, getUser.Bio)
	assert.Equal(t, user.BirthDay, getUser.BirthDay)
	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.Avatar, getUser.Avatar)
	assert.Equal(t, user.Password, getUser.Password)
	assert.Equal(t, user.RefreshToken, getUser.RefreshToken)

	// Test Update User
	user.FirstName = "Test Oybek"
	user.LastName = "Test Atamatov"
	updUser, err := repo.UpdateUserDB(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, updUser.Id)
	assert.Equal(t, user.FirstName, updUser.FirstName)
	assert.Equal(t, user.LastName, updUser.LastName)
	assert.Equal(t, user.Username, updUser.Username)
	assert.Equal(t, user.Bio, updUser.Bio)
	assert.Equal(t, user.Email, updUser.Email)

	// Test Delete User
	err = repo.DeleteUserDB(&pbu.IdRequest{Id: user.Id} )
	assert.NoError(t, err)

	// Test List Users
	Users, err := repo.ListUserDB(&pbu.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, Users)
}