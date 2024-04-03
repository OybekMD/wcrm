package repo

import (
	pbc "post-service/genproto/post"
)

// CouserStorageI ...
type PostStorageI interface {
	CreateCategoryIconDB(*pbc.CategoryIcon) (*pbc.CategoryIcon, error)
	ReadCategoryIconDB(*pbc.IdRequest) (*pbc.CategoryIcon, error)
	UpdateCategoryIconDB(*pbc.CategoryIcon) (*pbc.CategoryIcon, error)
	DeleteCategoryIconDB(*pbc.IdRequest) (*pbc.MessageResponse, error)
	ListCategoryIconsDB(*pbc.GetAllRequest) (*pbc.ListCategoryIconResponse, error)

	CreateCategoryDB(*pbc.Category) (*pbc.Category, error)
	ReadCategoryDB(*pbc.IdRequest) (*pbc.Category, error)
	UpdateCategoryDB(*pbc.Category) (*pbc.Category, error)
	DeleteCategoryDB(*pbc.IdRequest) (*pbc.MessageResponse, error)
	ListCategorysDB(*pbc.GetAllRequest) (*pbc.ListCategoryResponse, error)

	CreateProductDB(*pbc.Product) (*pbc.Product, error)
	ReadProductDB(*pbc.IdRequest) (*pbc.Product, error)
	UpdateProductDB(*pbc.Product) (*pbc.Product, error)
	DeleteProductDB(*pbc.IdRequest) (*pbc.MessageResponse, error)
	ListProductsDB(*pbc.GetAllRequest) (*pbc.ListProductResponse, error)
	ListProductsWithCommentsDB(*pbc.GetAllRequest) (*pbc.ListProductWithCommentResponse, error)

	CreateOrderproductDB(*pbc.Orderproduct) (*pbc.Orderproduct, error)
	ReadOrderproductDB(*pbc.IdRequest) (*pbc.Orderproduct, error)
	UpdateOrderproductDB(*pbc.Orderproduct) (*pbc.Orderproduct, error)
	DeleteOrderproductDB(*pbc.IdRequest) (*pbc.MessageResponse, error)
	ListOrderproductsDB(*pbc.GetAllRequest) (*pbc.ListOrderproductResponse, error)
}
