package handlers

import (
	"api-gateway/api/api_test/storage/kv"
	pbu "api-gateway/genproto/user"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/k0kubun/pp"
)

func CreateUser(c *gin.Context) {
	var newUser pbu.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	userString, err := kv.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp pbu.User
	if err := json.Unmarshal([]byte(userString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if err := kv.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func ListUsers(c *gin.Context) {
	usersStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	pp.Println(usersStrings)

	var users []*pbu.User
	for _, userString := range usersStrings {
		var user pbu.User
		if err := json.Unmarshal([]byte(userString), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}