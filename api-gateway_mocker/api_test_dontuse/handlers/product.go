package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbp "api-gateway/genproto/post"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Product crud
func CreateProduct(c *gin.Context) {
	var newProduct pbp.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productJson, err := json.Marshal(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newProduct.Id), string(productJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func GetProduct(c *gin.Context) {
	productID := c.Query("id")
	productString, err := kv.Get(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbp.Product
	if err := json.Unmarshal([]byte(productString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteProduct(c *gin.Context) {
	productId := c.Query("id")
	if err := kv.Delete(productId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product was deleted successfully",
	})
}

func ListProducts(c *gin.Context) {
	productsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var products []*pbp.Product
	for _, productString := range productsStrings {
		var product pbp.Product
		if err := json.Unmarshal([]byte(productString), &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		products = append(products, &product)
	}

	c.JSON(http.StatusOK, products)
}
