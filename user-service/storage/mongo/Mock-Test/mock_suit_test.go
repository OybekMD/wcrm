package mongo

import (
	"testing"

	"user-service/config"
	pbu "user-service/genproto/user"
	"user-service/pkg/db"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUserMockMongoDB(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err := db.ConnectToMongoDB(cfg)
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
	readUser, err := repo.ReadUserDB(&pbu.IdRequest{Id: user.Id})
	assert.NoError(t, err)
	assert.Equal(t, user.Id, readUser.Id)
	assert.Equal(t, user.FirstName, readUser.FirstName)
	assert.Equal(t, user.LastName, readUser.LastName)
	assert.Equal(t, user.Username, readUser.Username)
	assert.Equal(t, user.PhoneNumber, readUser.PhoneNumber)
	assert.Equal(t, user.Bio, readUser.Bio)
	assert.NotNil(t, user.BirthDay, readUser.BirthDay)
	assert.Equal(t, user.Email, readUser.Email)
	assert.Equal(t, user.Avatar, readUser.Avatar)

	// Test Update User
	user.FirstName = "Test"
	user.LastName = "Test"
	updUser, err := repo.UpdateUserDB(user)
	assert.NoError(t, err)
	assert.Equal(t, user.Id, updUser.Id)
	assert.Equal(t, user.FirstName, updUser.FirstName)
	assert.Equal(t, user.LastName, updUser.LastName)
	assert.Equal(t, user.Username, updUser.Username)
	assert.Equal(t, user.PhoneNumber, createdUser.PhoneNumber)
	assert.Equal(t, user.Bio, updUser.Bio)
	assert.NotNil(t, user.BirthDay, createdUser.BirthDay)
	assert.Equal(t, user.Email, updUser.Email)
	assert.Equal(t, user.Avatar, updUser.Avatar)

	// Test List Users
	Users, err := repo.ListUserDB(&pbu.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, Users)

	// Test Delete User
	message, err := repo.DeleteUserDB(&pbu.IdRequest{Id: user.Id})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "User successfully deleted!")
}
