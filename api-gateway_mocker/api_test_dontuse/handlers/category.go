package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbp "api-gateway/genproto/post"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Category crud
func CreateCategory(c *gin.Context) {
	var newCategory pbp.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	categoryJson, err := json.Marshal(newCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newCategory.Id), string(categoryJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newCategory)
}

func GetCategory(c *gin.Context) {
	categoryID := c.Query("id")
	categoryString, err := kv.Get(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbp.Category
	if err := json.Unmarshal([]byte(categoryString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteCategory(c *gin.Context) {
	categoryId := c.Query("id")
	if err := kv.Delete(categoryId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category was deleted successfully",
	})
}

func ListCategorys(c *gin.Context) {
	categorysStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var categorys []*pbp.Category
	for _, categoryString := range categorysStrings {
		var category pbp.Category
		if err := json.Unmarshal([]byte(categoryString), &category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		categorys = append(categorys, &category)
	}

	c.JSON(http.StatusOK, categorys)
}
