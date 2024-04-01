package mongo

import (
	"strconv"
	"testing"

	"post-service/config"
	pbp "post-service/genproto/post"
	"post-service/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestCategoryIconMongoDB(t *testing.T) {
    // Connect to database
    cfg := config.Load()
    db, err := db.ConnectToMongoDB(cfg)
    if err != nil {
        t.Fatal(err)
    }

    // Test Create CategoryIcon
    repo := NewPostMongoRepo(db)
    categoryIcon := &pbp.CategoryIcon{
        Name:    "TestCategory",
        Picture: "example.com/test.jpg",
    }


    createdCategoryIcon, err := repo.CreateCategoryIconDB(categoryIcon)
	id := strconv.Itoa(int(createdCategoryIcon.Id))
    assert.NoError(t, err)
    assert.Equal(t, categoryIcon.Id, createdCategoryIcon.Id)
    assert.Equal(t, categoryIcon.Name, createdCategoryIcon.Name)
    assert.Equal(t, categoryIcon.Picture, createdCategoryIcon.Picture)

    // Test Read CategoryIcon
    readCategoryIcon, err := repo.ReadCategoryIconDB(&pbp.IdRequest{Id: id})
    assert.NoError(t, err)
    assert.Equal(t, categoryIcon.Id, readCategoryIcon.Id)
    assert.Equal(t, categoryIcon.Name, readCategoryIcon.Name)
    assert.Equal(t, categoryIcon.Picture, readCategoryIcon.Picture)

    // Test Update CategoryIcon
    categoryIcon.Name = "UpdatedCategory"
    categoryIcon.Picture = "example.com/updated.jpg"
    updatedCategoryIcon, err := repo.UpdateCategoryIconDB(categoryIcon)
    assert.NoError(t, err)
    assert.Equal(t, categoryIcon.Id, updatedCategoryIcon.Id)
    assert.Equal(t, categoryIcon.Name, updatedCategoryIcon.Name)
    assert.Equal(t, categoryIcon.Picture, updatedCategoryIcon.Picture)

    // Test List CategoryIcons
    categoryIcons, err := repo.ListCategoryIconsDB(&pbp.GetAllRequest{Page: 1, Limit: 10})
    assert.NoError(t, err)
    assert.NotEmpty(t, categoryIcons)

    // Test Delete CategoryIcon
    message, err := repo.DeleteCategoryIconDB(&pbp.IdRequest{Id: id})
    assert.NoError(t, err)
    assert.Equal(t, message.Message, "CategoryIcon successfully deleted!")
}
