package tests

import (
	"api-gateway/test_api/handler"
	"api-gateway/test_api/storage"
	"encoding/json"
	"io"
	"strconv"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {

	gin.SetMode(gin.TestMode)

	// USER
	require.NoError(t, SetupMinimumInstance(""))
	file, err := OpenFile("user.json")

	require.NoError(t, err)
	req := NewRequest(http.MethodPost, "/user", file)
	res := httptest.NewRecorder()

	r := gin.Default()
	r.POST("/user", handler.CreateUser)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var user *storage.User

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	require.Equal(t, user.FirstName, "John")
	require.Equal(t, user.LastName, "Doe")
	require.Equal(t, user.Username, "johndoe")
	require.Equal(t, user.PhoneNumber, "+1234567890")
	require.Equal(t, user.Bio, "Software Engineer")
	require.Equal(t, user.BirthDay, "1990-05-15")
	require.Equal(t, user.Email, "john@example.com")
	require.Equal(t, user.Avatar, "https://example.com/avatar.jpg")
	require.Equal(t, user.Password, "hashed_password_here")
	require.Equal(t, user.RefreshToken, "refresh_token_here")
	require.NotNil(t, user.CreatedAt)
	require.NotNil(t, user.UpdatedAt)

	// READ
	getReq := NewRequest(http.MethodGet, "/user/:id", nil)
	args := getReq.URL.Query()
	args.Add("id", user.Id)
	getReq.URL.RawQuery = args.Encode()
	getRes := httptest.NewRecorder()

	r.GET("/user/:id", handler.ReadUser)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)

	// Check User equal
	var getUser *storage.User
	bdByte, err := io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getUser))
	assert.Equal(t, user.Id, getUser.Id)
	assert.Equal(t, user.FirstName, getUser.FirstName)
	assert.Equal(t, user.LastName, getUser.LastName)
	assert.Equal(t, user.Username, getUser.Username)
	assert.Equal(t, user.PhoneNumber, getUser.PhoneNumber)
	assert.Equal(t, user.Bio, getUser.Bio)
	assert.Equal(t, user.BirthDay, getUser.BirthDay)
	assert.Equal(t, user.Email, getUser.Email)
	assert.Equal(t, user.Avatar, getUser.Avatar)
	assert.Equal(t, user.Password, getUser.Password)
	assert.Equal(t, user.RefreshToken, getUser.RefreshToken)
	assert.Equal(t, user.CreatedAt, getUser.CreatedAt)
	assert.Equal(t, user.UpdatedAt, getUser.UpdatedAt)

	// GetALL or List
	ListReq := NewRequest(http.MethodGet, "/users", file)
	listRes := httptest.NewRecorder()

	r.GET("/users", handler.GetAllUsers)
	r.ServeHTTP(listRes, ListReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bLists, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq := NewRequest(http.MethodDelete, "/user/:id?id="+user.Id, file)
	delRes := httptest.NewRecorder()

	r.DELETE("/user/:id", handler.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resUserB storage.ResponseMessage
	bDel, err := io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resUserB))

	// --- CategoryIcon ---
	fileCategoryIcon, err := OpenFile("categoryicon.json")

	require.NoError(t, err)
	req = NewRequest(http.MethodPost, "/categoryicon", fileCategoryIcon)
	res = httptest.NewRecorder()
	r = gin.Default()

	r.POST("/categoryicon", handler.CreateCategoryIcon)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var categoryicon storage.CategoryIcon

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &categoryicon))
	require.Equal(t, categoryicon.Id, int64(1))
	require.Equal(t, categoryicon.Name, "Food")
	require.Equal(t, categoryicon.Picture, "https://example.com/food_icon.png")

	// READ
	getReqCategoryIcon := NewRequest(http.MethodGet, "/categoryicon/:id", fileCategoryIcon)

	argsCategoryIcon := getReqCategoryIcon.URL.Query()
	idci := strconv.Itoa(int(categoryicon.Id))
	argsCategoryIcon.Add("id", idci)
	getReqCategoryIcon.URL.RawQuery = argsCategoryIcon.Encode()
	getResCategoryIcon := httptest.NewRecorder()

	r = gin.Default()
	r.GET("/categoryicon/:id", handler.ReadCategoryIcon)
	r.ServeHTTP(getResCategoryIcon, getReqCategoryIcon)
	assert.Equal(t, http.StatusOK, getResCategoryIcon.Code)

	// Check User equal
	var getCategoryIcon *storage.CategoryIcon
	bdByte, err = io.ReadAll(getResCategoryIcon.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getCategoryIcon))
	assert.Equal(t, categoryicon.Id, getCategoryIcon.Id)
	assert.Equal(t, categoryicon.Name, getCategoryIcon.Name)
	assert.Equal(t, categoryicon.Picture, getCategoryIcon.Picture)

	// GetALL or List
	ListReqCategoryIcon := NewRequest(http.MethodGet, "/categoryicons", file)
	listResCategoryIcon := httptest.NewRecorder()

	r.GET("/categoryicons", handler.ListCategoryIcons)
	r.ServeHTTP(listResCategoryIcon, ListReqCategoryIcon)
	assert.Equal(t, http.StatusOK, listResCategoryIcon.Code)
	bLists, err = io.ReadAll(listResCategoryIcon.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq = NewRequest(http.MethodDelete, "/categoryicon/:id?id="+idci, file)
	delRes = httptest.NewRecorder()

	r.DELETE("/categoryicon/:id", handler.DeleteCategoryIcon)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resCategoryIconB storage.ResponseMessage
	bDel, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resCategoryIconB))

	// --- Category ---
	fileCategory, err := OpenFile("category.json")

	require.NoError(t, err)
	req = NewRequest(http.MethodPost, "/category", fileCategory)
	res = httptest.NewRecorder()
	r = gin.Default()

	r.POST("/category", handler.CreateCategory)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var category storage.Category

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &category))
	require.Equal(t, category.Id, int64(1))
	require.Equal(t, category.Name, "Food & Drinks")
	require.Equal(t, category.IconId, int64(1))
	require.NotNil(t, category.CreatedAt)
	require.NotNil(t, category.UpdatedAt)

	// READ
	getReqCategory := NewRequest(http.MethodGet, "/category/:id", fileCategory)

	argsCategory := getReqCategory.URL.Query()
	idc := strconv.Itoa(int(category.Id))
	argsCategory.Add("id", idc)
	getReqCategory.URL.RawQuery = argsCategory.Encode()
	getResCategory := httptest.NewRecorder()

	r = gin.Default()
	r.GET("/category/:id", handler.ReadCategory)
	r.ServeHTTP(getResCategory, getReqCategory)
	assert.Equal(t, http.StatusOK, getResCategory.Code)

	// Check User equal
	var getCategory *storage.Category
	bdByte, err = io.ReadAll(getResCategory.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getCategory))
	assert.Equal(t, category.Id, getCategory.Id)
	assert.Equal(t, category.Name, getCategory.Name)
	assert.Equal(t, category.IconId, getCategory.IconId)
	assert.Equal(t, category.CreatedAt, getCategory.CreatedAt)
	assert.Equal(t, category.UpdatedAt, getCategory.UpdatedAt)

	// GetALL or List
	ListReqCategory := NewRequest(http.MethodGet, "/categorys", file)
	listResCategory := httptest.NewRecorder()

	r.GET("/categorys", handler.ListCategorys)
	r.ServeHTTP(listResCategory, ListReqCategory)
	assert.Equal(t, http.StatusOK, listResCategory.Code)
	bLists, err = io.ReadAll(listResCategory.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq = NewRequest(http.MethodDelete, "/category/:id?id="+idc, file)
	delRes = httptest.NewRecorder()

	r.DELETE("/category/:id", handler.DeleteCategory)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resCategoryB storage.ResponseMessage
	bDel, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resCategoryB))

	// --- Product ---
	fileProduct, err := OpenFile("product.json")

	require.NoError(t, err)
	req = NewRequest(http.MethodPost, "/product", fileProduct)
	res = httptest.NewRecorder()
	r = gin.Default()

	r.POST("/product", handler.CreateProduct)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var product storage.Product

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &product))
	require.Equal(t, product.Id, int64(1))
	require.Equal(t, product.Title, "Gammburger")
	require.Equal(t, product.Description, "Fresh Gammburger, rich in healthy fats.")
	require.Equal(t, product.Price, int64(30000))
	require.Equal(t, product.Picture, "https://example.com/gammburger.jpg")
	require.Equal(t, product.CategoryId, int64(1))
	require.NotNil(t, product.CreatedAt)
	require.NotNil(t, product.UpdatedAt)

	// READ
	getReqProduct := NewRequest(http.MethodGet, "/product/:id", fileProduct)

	argsProduct := getReqProduct.URL.Query()
	idp := strconv.Itoa(int(product.Id))
	argsProduct.Add("id", idp)
	getReqProduct.URL.RawQuery = argsProduct.Encode()
	getResProduct := httptest.NewRecorder()

	r = gin.Default()
	r.GET("/product/:id", handler.ReadProduct)
	r.ServeHTTP(getResProduct, getReqProduct)
	assert.Equal(t, http.StatusOK, getResProduct.Code)

	// Check User equal
	var getProduct *storage.Product
	bdByte, err = io.ReadAll(getResProduct.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getProduct))
	assert.Equal(t, product.Id, getProduct.Id)
	assert.Equal(t, product.Title, getProduct.Title)
	assert.Equal(t, product.Description, getProduct.Description)
	assert.Equal(t, product.Picture, getProduct.Picture)
	assert.Equal(t, product.CategoryId, getProduct.CategoryId)
	assert.Equal(t, product.CreatedAt, getProduct.CreatedAt)
	assert.Equal(t, product.UpdatedAt, getProduct.UpdatedAt)

	// GetALL or List
	ListReqProduct := NewRequest(http.MethodGet, "/products", file)
	listResProduct := httptest.NewRecorder()

	r.GET("/products", handler.ListProducts)
	r.ServeHTTP(listResProduct, ListReqProduct)
	assert.Equal(t, http.StatusOK, listResProduct.Code)
	bLists, err = io.ReadAll(listResProduct.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq = NewRequest(http.MethodDelete, "/product/:id?id="+idp, file)
	delRes = httptest.NewRecorder()

	r.DELETE("/product/:id", handler.DeleteProduct)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resProductB storage.ResponseMessage
	bDel, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resProductB))

	// --- Orderproduct ---
	fileOrderproduct, err := OpenFile("orderproduct.json")

	require.NoError(t, err)
	req = NewRequest(http.MethodPost, "/orderproduct", fileOrderproduct)
	res = httptest.NewRecorder()
	r = gin.Default()

	r.POST("/orderproduct", handler.CreateOrderproduct)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var orderproduct storage.Orderproduct

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &orderproduct))
	require.Equal(t, orderproduct.Id, int64(1))
	require.Equal(t, orderproduct.UserId, "f7d14b9d-6a18-4b11-bf7c-78eb810e8c7f")
	require.Equal(t, orderproduct.ProductId, int64(1))
	require.NotNil(t, orderproduct.CreatedAt)

	// READ
	getReqOrderproduct := NewRequest(http.MethodGet, "/orderproduct/:id", fileOrderproduct)

	argsOrderproduct := getReqOrderproduct.URL.Query()
	idop := strconv.Itoa(int(orderproduct.Id))
	argsOrderproduct.Add("id", idop)
	getReqOrderproduct.URL.RawQuery = argsOrderproduct.Encode()
	getResOrderproduct := httptest.NewRecorder()

	r = gin.Default()
	r.GET("/orderproduct/:id", handler.ReadOrderproduct)
	r.ServeHTTP(getResOrderproduct, getReqOrderproduct)
	assert.Equal(t, http.StatusOK, getResOrderproduct.Code)

	// Check User equal
	var getOrderproduct *storage.Orderproduct
	bdByte, err = io.ReadAll(getResOrderproduct.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getOrderproduct))
	assert.Equal(t, orderproduct.Id, getOrderproduct.Id)
	assert.Equal(t, orderproduct.UserId, getOrderproduct.UserId)
	assert.Equal(t, orderproduct.ProductId, getOrderproduct.ProductId)
	assert.Equal(t, orderproduct.CreatedAt, getOrderproduct.CreatedAt)

	// GetALL or List
	ListReqOrderproduct := NewRequest(http.MethodGet, "/orderproducts", file)
	listResOrderproduct := httptest.NewRecorder()

	r.GET("/orderproducts", handler.ListOrderproducts)
	r.ServeHTTP(listResOrderproduct, ListReqOrderproduct)
	assert.Equal(t, http.StatusOK, listResOrderproduct.Code)
	bLists, err = io.ReadAll(listResOrderproduct.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq = NewRequest(http.MethodDelete, "/orderproduct/:id?id="+idop, file)
	delRes = httptest.NewRecorder()

	r.DELETE("/orderproduct/:id", handler.DeleteOrderproduct)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resOrderproductB storage.ResponseMessage
	bDel, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resOrderproductB))

	// --- Comment ---
	fileComment, err := OpenFile("comment.json")

	require.NoError(t, err)
	req = NewRequest(http.MethodPost, "/comment", fileComment)
	res = httptest.NewRecorder()
	r = gin.Default()

	r.POST("/comment", handler.CreateComment)
	r.ServeHTTP(res, req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.Code)

	var comment storage.Comment

	// CREATE
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &comment))
	require.Equal(t, comment.Id, int64(1))
	require.Equal(t, comment.UserId, "f7d14b9d-6a18-4b11-bf7c-78eb810e8c7f")
	require.Equal(t, comment.ProductId, int64(1))
	require.NotNil(t, comment.CreatedAt)

	// READ
	getReqComment := NewRequest(http.MethodGet, "/comment/:id", fileComment)

	argsComment := getReqComment.URL.Query()
	idcomment := strconv.Itoa(int(comment.Id))
	argsComment.Add("id", idcomment)
	getReqComment.URL.RawQuery = argsComment.Encode()
	getResComment := httptest.NewRecorder()

	r = gin.Default()
	r.GET("/comment/:id", handler.ReadComment)
	r.ServeHTTP(getResComment, getReqComment)
	assert.Equal(t, http.StatusOK, getResComment.Code)

	// Check User equal
	var getComment *storage.Comment
	bdByte, err = io.ReadAll(getResComment.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bdByte, &getComment))
	assert.Equal(t, comment.Id, getComment.Id)
	assert.Equal(t, comment.UserId, getComment.UserId)
	assert.Equal(t, comment.ProductId, getComment.ProductId)
	assert.Equal(t, comment.CreatedAt, getComment.CreatedAt)

	// GetALL or List
	ListReqComment := NewRequest(http.MethodGet, "/comments", file)
	listResComment := httptest.NewRecorder()

	r.GET("/comments", handler.ListComments)
	r.ServeHTTP(listResComment, ListReqComment)
	assert.Equal(t, http.StatusOK, listResComment.Code)
	bLists, err = io.ReadAll(listResComment.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bLists)

	// DELETE
	delReq = NewRequest(http.MethodDelete, "/comment/:id?id="+idcomment, file)
	delRes = httptest.NewRecorder()

	r.DELETE("/comment/:id", handler.DeleteComment)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)

	var resCommentB storage.ResponseMessage
	bDel, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bDel, &resCommentB))

}
