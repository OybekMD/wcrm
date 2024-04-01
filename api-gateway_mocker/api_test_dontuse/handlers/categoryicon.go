package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbp "api-gateway/genproto/post"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// CategoryIcon crud
func CreateCategoryIcon(c *gin.Context) {
	var newCategoryIcon pbp.CategoryIcon
	if err := c.ShouldBindJSON(&newCategoryIcon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	categoryIconJson, err := json.Marshal(newCategoryIcon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newCategoryIcon.Id), string(categoryIconJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newCategoryIcon)
}

func GetCategoryIcon(c *gin.Context) {
	categoryIconID := c.Query("id")
	categoryIconString, err := kv.Get(categoryIconID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbp.CategoryIcon
	if err := json.Unmarshal([]byte(categoryIconString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteCategoryIcon(c *gin.Context) {
	categoryIconId := c.Query("id")
	if err := kv.Delete(categoryIconId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "CategoryIcon was deleted successfully",
	})
}

func ListCategoryIcons(c *gin.Context) {
	categoryIconsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var categoryIcons []*pbp.CategoryIcon
	for _, categoryIconString := range categoryIconsStrings {
		var categoryIcon pbp.CategoryIcon
		if err := json.Unmarshal([]byte(categoryIconString), &categoryIcon); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		categoryIcons = append(categoryIcons, &categoryIcon)
	}

	c.JSON(http.StatusOK, categoryIcons)
}
