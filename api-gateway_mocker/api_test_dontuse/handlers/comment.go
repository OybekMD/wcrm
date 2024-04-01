package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbp "api-gateway/genproto/comment"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Comment crud
func CreateComment(c *gin.Context) {
	var newComment pbp.Comment
	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	commentJson, err := json.Marshal(newComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newComment.Id), string(commentJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newComment)
}

func GetComment(c *gin.Context) {
	commentID := c.Query("id")
	commentString, err := kv.Get(commentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbp.Comment
	if err := json.Unmarshal([]byte(commentString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteComment(c *gin.Context) {
	commentId := c.Query("id")
	if err := kv.Delete(commentId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Comment was deleted successfully",
	})
}

func ListComments(c *gin.Context) {
	commentsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var comments []*pbp.Comment
	for _, commentString := range commentsStrings {
		var comment pbp.Comment
		if err := json.Unmarshal([]byte(commentString), &comment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		comments = append(comments, &comment)
	}

	c.JSON(http.StatusOK, comments)
}
