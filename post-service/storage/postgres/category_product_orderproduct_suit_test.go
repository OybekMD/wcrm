package postgres

import (
	"strconv"
	"testing"

	"post-service/config"
	pbp "post-service/genproto/post"
	"post-service/pkg/db"

	"github.com/stretchr/testify/assert"
)

func TestCategoryPostgres(t *testing.T) {
	// Connect to database
	cfg := config.Load()
	db, err := db.ConnectToDB(cfg)
	if err != nil {
		return
	}

	repo := NewPostRepo(db)
	category := &pbp.Category{
		Name:   "Gamburgers",
		IconId: 1,
	}

	product := &pbp.Product{
		Title:       "Do'mboq Gamburger",
		Description: "Gamburger tami tiligizda ajoyib hissiyot qildirajak",
		Price:       17000,
		Picture:     "example.com/gamburger.jpg",
	}
	orderproduct := &pbp.Orderproduct{
		UserId: "67686d61-45a7-4a88-9f61-6a869a3b97aa",
	}

	// Test Create Category
	createdCategory, err := repo.CreateCategoryDB(category)
	category.Id = createdCategory.Id
	idc := strconv.Itoa(int(createdCategory.Id))
	assert.NoError(t, err)
	assert.NotNil(t, createdCategory.Id)
	assert.Equal(t, category.Name, createdCategory.Name)
	assert.Equal(t, category.IconId, createdCategory.IconId)
	assert.NotNil(t, createdCategory.CreatedAt)
	assert.NotNil(t, createdCategory.UpdatedAt)

	// Test Create Product
	product.CategoryId = createdCategory.Id
	createdProduct, err := repo.CreateProductDB(product)
	product.Id = createdProduct.Id
	idp := strconv.Itoa(int(createdProduct.Id))
	assert.NoError(t, err)
	assert.NotNil(t, createdProduct.Id)
	assert.Equal(t, product.Title, createdProduct.Title)
	assert.Equal(t, product.Description, createdProduct.Description)
	assert.Equal(t, product.Price, createdProduct.Price)
	assert.Equal(t, product.Picture, createdProduct.Picture)
	assert.Equal(t, product.CategoryId, createdProduct.CategoryId)
	assert.NotNil(t, createdProduct.CreatedAt)
	assert.NotNil(t, createdProduct.UpdatedAt)

	// Test Create Orderproduct
	orderproduct.ProductId = createdProduct.Id
	createdOrderproduct, err := repo.CreateOrderproductDB(orderproduct)
	orderproduct.Id = createdOrderproduct.Id
	ido := strconv.Itoa(int(createdOrderproduct.Id))
	assert.NoError(t, err)
	assert.NotNil(t, createdOrderproduct.Id)
	assert.Equal(t, orderproduct.UserId, createdOrderproduct.UserId)
	assert.Equal(t, orderproduct.ProductId, createdOrderproduct.ProductId)
	assert.NotNil(t, createdProduct.CreatedAt)

	//Test Read Category
	idCategory, err := repo.ReadCategoryDB(&pbp.IdRequest{Id: idc})
	assert.NoError(t, err)
	assert.Equal(t, category.Id, idCategory.Id)
	assert.Equal(t, category.Name, idCategory.Name)
	assert.Equal(t, category.IconId, idCategory.IconId)

	//Test Read Product
	idProduct, err := repo.ReadProductDB(&pbp.IdRequest{Id: idp})
	assert.NoError(t, err)
	assert.Equal(t, product.Id, idProduct.Id)
	assert.Equal(t, product.Title, idProduct.Title)
	assert.Equal(t, product.Description, idProduct.Description)
	assert.Equal(t, product.Price, idProduct.Price)
	assert.Equal(t, product.Picture, idProduct.Picture)
	assert.Equal(t, product.CategoryId, idProduct.CategoryId)

	//Test Read Orderproduct
	idOrderproduct, err := repo.ReadOrderproductDB(&pbp.IdRequest{Id: ido})
	assert.NoError(t, err)
	assert.Equal(t, orderproduct.Id, idOrderproduct.Id)
	assert.Equal(t, orderproduct.UserId, idOrderproduct.UserId)
	assert.Equal(t, orderproduct.ProductId, idOrderproduct.ProductId)

	// Test Update Category
	category.Name = "Update"
	category.IconId = 2
	updCategory, err := repo.UpdateCategoryDB(category)
	assert.NoError(t, err)
	assert.Equal(t, category.Id, updCategory.Id)
	assert.Equal(t, category.Name, updCategory.Name)
	assert.Equal(t, category.IconId, updCategory.IconId)

	// Test Update Product
	product.Title = "Update"
	product.Description = "upd description"
	product.Price = 20000
	product.Picture = "update"
	product.CategoryId = createdCategory.Id
	updProduct, err := repo.UpdateProductDB(product)
	assert.NoError(t, err)
	assert.Equal(t, product.Id, updProduct.Id)
	assert.Equal(t, product.Title, updProduct.Title)
	assert.Equal(t, product.Description, updProduct.Description)
	assert.Equal(t, product.Price, updProduct.Price)
	assert.Equal(t, product.Picture, updProduct.Picture)
	assert.Equal(t, product.CategoryId, updProduct.CategoryId)

	// Test Update Orderproduct
	orderproduct.UserId = "7217372f-a31e-4403-b7a3-0cc963af589f"
	orderproduct.ProductId = createdProduct.Id
	updOrderproduct, err := repo.UpdateOrderproductDB(orderproduct)
	assert.NoError(t, err)
	assert.Equal(t, orderproduct.Id, updOrderproduct.Id)
	assert.Equal(t, orderproduct.UserId, updOrderproduct.UserId)
	assert.Equal(t, orderproduct.ProductId, updOrderproduct.ProductId)

	// Test List Categorys
	categorys, err := repo.ListCategorysDB(&pbp.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, categorys)

	// Test List Products
	products, err := repo.ListProductsDB(&pbp.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, products)

	// Test List Orderproduct
	orderproducts, err := repo.ListOrderproductsDB(&pbp.GetAllRequest{Page: 1, Limit: 10})
	assert.NoError(t, err)
	assert.NotEmpty(t, orderproducts)

	// Test Delete Category
	message, err := repo.DeleteCategoryDB(&pbp.IdRequest{Id: idc})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "Category successfully deleted!")

	// Test Delete Product
	message, err = repo.DeleteProductDB(&pbp.IdRequest{Id: idp})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "Product successfully deleted!")

	// Test Delete Orderproduct
	message, err = repo.DeleteOrderproductDB(&pbp.IdRequest{Id: ido})
	assert.NoError(t, err)
	assert.Equal(t, message.Message, "Orderproduct successfully deleted!")
}
