package handler

import (
	"api-gateway/test_api/storage"
	"api-gateway/test_api/storage/kv"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RegisterUser(c *gin.Context) {
	var newUser storage.RegisterModel
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return
	}

	if err := kv.Set(id.String(), string(userJson)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "We send verification password you email",
	})
}

func Verification(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func CreateUser(c *gin.Context) {
	var newUser storage.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJ, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJ)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func ReadUser(c *gin.Context) {
	id := c.Query("id")

	userGet, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.User
	if err := json.Unmarshal([]byte(userGet), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has been deleted",
	})
}

func GetAllUsers(c *gin.Context) {
	userList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*storage.User
	for _, l := range userList {
		var user storage.User

		if err := json.Unmarshal([]byte(l), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}

// -----------------------------------------------------------------

func CreateCategoryIcon(c *gin.Context) {
	var body storage.CategoryIcon

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = 1

	catei, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(body.Id))
	if err := kv.Set(id, string(catei)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func ReadCategoryIcon(c *gin.Context) {
	id := c.Query("id")

	userGet, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.CategoryIcon
	if err := json.Unmarshal([]byte(userGet), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListCategoryIcons(c *gin.Context) {
	prodList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var cis []*storage.CategoryIcon
	for _, l := range prodList {
		var ci storage.CategoryIcon

		if err := json.Unmarshal([]byte(l), &ci); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		cis = append(cis, &ci)
	}

	c.JSON(http.StatusOK, cis)
}

func DeleteCategoryIcon(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "CategoryIcon has been deleted",
	})
}

// -----------------------------------------------------------------

func CreateCategory(c *gin.Context) {
	var body storage.Category

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = 1

	cate, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(body.Id))
	if err := kv.Set(id, string(cate)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func ReadCategory(c *gin.Context) {
	id := c.Query("id")

	userGet, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.Category
	if err := json.Unmarshal([]byte(userGet), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListCategorys(c *gin.Context) {
	prodList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var categorys []*storage.Category
	for _, l := range prodList {
		var category storage.Category

		if err := json.Unmarshal([]byte(l), &category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		categorys = append(categorys, &category)
	}

	c.JSON(http.StatusOK, categorys)
}

func DeleteCategory(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "Category has been deleted",
	})
}

// -----------------------------------------------------------------

func CreateProduct(c *gin.Context) {
	var body storage.Product

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = 1

	item, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(body.Id))
	if err := kv.Set(id, string(item)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func ReadProduct(c *gin.Context) {
	id := c.Query("id")

	eget, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.Product
	if err := json.Unmarshal([]byte(eget), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListProducts(c *gin.Context) {
	eList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var items []*storage.Product
	for _, l := range eList {
		var item storage.Product

		if err := json.Unmarshal([]byte(l), &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		items = append(items, &item)
	}

	c.JSON(http.StatusOK, items)
}

func DeleteProduct(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "Category has been deleted",
	})
}

// -----------------------------------------------------------------

func CreateOrderproduct(c *gin.Context) {
	var body storage.Orderproduct

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = 1

	item, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(body.Id))
	if err := kv.Set(id, string(item)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func ReadOrderproduct(c *gin.Context) {
	id := c.Query("id")

	eget, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.Orderproduct
	if err := json.Unmarshal([]byte(eget), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListOrderproducts(c *gin.Context) {
	eList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var items []*storage.Orderproduct
	for _, l := range eList {
		var item storage.Orderproduct

		if err := json.Unmarshal([]byte(l), &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		items = append(items, &item)
	}

	c.JSON(http.StatusOK, items)
}

func DeleteOrderproduct(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "Category has been deleted",
	})
}

// -----------------------------------------------------------------

func CreateComment(c *gin.Context) {
	var body storage.Comment

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	body.Id = 1

	item, err := json.Marshal(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := strconv.Itoa(int(body.Id))
	if err := kv.Set(id, string(item)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, body)
}

func ReadComment(c *gin.Context) {
	id := c.Query("id")

	eget, err := kv.Get(id)
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var res storage.Comment
	if err := json.Unmarshal([]byte(eget), &res); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func ListComments(c *gin.Context) {
	eList, err := kv.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var items []*storage.Comment
	for _, l := range eList {
		var item storage.Comment

		if err := json.Unmarshal([]byte(l), &item); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		items = append(items, &item)
	}

	c.JSON(http.StatusOK, items)
}

func DeleteComment(c *gin.Context) {
	id := c.Query("id")

	if err := kv.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"content": "Category has been deleted",
	})
}