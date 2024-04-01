package postgres

import (
	"strconv"
	"testing"

	"comment-service/config"
	pbc "comment-service/genproto/comment"
	"comment-service/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestCommentPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	// Test Create Comment
	repo := NewCommentRepo(db)
	comment := &pbc.Comment{
		Content:   "This is amazing",
		UserId:    "30b9eb83-edee-4735-a976-3ce489a32190",
		ProductId: 1,
	}

	createdComment, err := repo.CreateCommentDB(comment)
	comment.Id = createdComment.Id
	id := strconv.Itoa(int(createdComment.Id))
	assert.NoError(t, err)
	assert.NotNil(t, createdComment.Id)
	assert.Equal(t, comment.Content, createdComment.Content)
	assert.Equal(t, comment.UserId, createdComment.UserId)
	assert.Equal(t, comment.ProductId, createdComment.ProductId)
	assert.NotNil(t, createdComment.CreatedAt)
	assert.NotNil(t, createdComment.UpdatedAt)

	//Test Read Comment
	idComment, err := repo.ReadCommentDB(&pbc.IdRequest{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, comment.Id, idComment.Id)
	assert.Equal(t, comment.Content, idComment.Content)
	assert.Equal(t, comment.UserId, idComment.UserId)
	assert.Equal(t, comment.ProductId, idComment.ProductId)

	// Test Update Comment
	comment.Content = "Update Content"
	updComment, err := repo.UpdateCommentDB(comment)
	assert.NoError(t, err)
	assert.Equal(t, comment.Id, updComment.Id)
	assert.Equal(t, comment.Content, updComment.Content)
	assert.Equal(t, comment.UserId, updComment.UserId)
	assert.Equal(t, comment.ProductId, updComment.ProductId)

	// Test List Comments
	comments, err := repo.ListCommentsDB(&pbc.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, comments)

	// Test Delete Comment
	message, err := repo.DeleteCommentDB(&pbc.IdRequest{Id: id})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "Comment successfully deleted!")
}
