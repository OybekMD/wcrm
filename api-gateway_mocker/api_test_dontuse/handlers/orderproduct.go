package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbp "api-gateway/genproto/post"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

// Orderproduct crud
func CreateOrderproduct(c *gin.Context) {
	var newOrderproduct pbp.Orderproduct
	if err := c.ShouldBindJSON(&newOrderproduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	orderproductJson, err := json.Marshal(newOrderproduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(cast.ToString(newOrderproduct.Id), string(orderproductJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newOrderproduct)
}

func GetOrderproduct(c *gin.Context) {
	orderproductID := c.Query("id")
	orderproductString, err := kv.Get(orderproductID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbp.Orderproduct
	if err := json.Unmarshal([]byte(orderproductString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteOrderproduct(c *gin.Context) {
	orderproductId := c.Query("id")
	if err := kv.Delete(orderproductId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orderproduct was deleted successfully",
	})
}

func ListOrderproducts(c *gin.Context) {
	orderproductsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var orderproducts []*pbp.Orderproduct
	for _, orderproductString := range orderproductsStrings {
		var orderproduct pbp.Orderproduct
		if err := json.Unmarshal([]byte(orderproductString), &orderproduct); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		orderproducts = append(orderproducts, &orderproduct)
	}

	c.JSON(http.StatusOK, orderproducts)
}
