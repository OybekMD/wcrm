package postgres

import (
	"strconv"
	"testing"

	"post-service/config"
	pbp "post-service/genproto/post"
	"post-service/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIconPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	// Test Create CategoryIcon
	repo := NewPostRepo(db)
	categoryIcon := &pbp.CategoryIcon{
		Name:    "Fast Food",
		Picture: "example.com/food.jpg",
	}

	createdCategoryIcon, err := repo.CreateCategoryIconDB(categoryIcon)
	categoryIcon.Id = createdCategoryIcon.Id
	id := strconv.Itoa(int(createdCategoryIcon.Id)) 
	assert.NoError(t, err)
	assert.NotNil(t, createdCategoryIcon.Id)
	assert.Equal(t, categoryIcon.Name, createdCategoryIcon.Name)
	assert.Equal(t, categoryIcon.Picture, createdCategoryIcon.Picture)

	//Test Read CategoryIcon
	idCategoryIcon, err := repo.ReadCategoryIconDB(&pbp.IdRequest{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, categoryIcon.Id, idCategoryIcon.Id)
	assert.Equal(t, categoryIcon.Name, idCategoryIcon.Name)
	assert.Equal(t, categoryIcon.Picture, idCategoryIcon.Picture)

	// Test Update CategoryIcon
	categoryIcon.Name = "Update"
	updCategoryIcon, err := repo.UpdateCategoryIconDB(categoryIcon)
	assert.NoError(t, err)
	assert.Equal(t, categoryIcon.Id, updCategoryIcon.Id)
	assert.Equal(t, categoryIcon.Name, updCategoryIcon.Name)
	assert.Equal(t, categoryIcon.Picture, updCategoryIcon.Picture)

	// Test List CategoryIcons
	categoryIcons, err := repo.ListCategoryIconsDB(&pbp.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, categoryIcons)

	// Test Delete CategoryIcon
	message, err := repo.DeleteCategoryIconDB(&pbp.IdRequest{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "CategoryIcon successfully deleted!")
}
