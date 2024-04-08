package service

import (
	"context"
	pbc "post-service/genproto/comment"
	pbp "post-service/genproto/post"
	l "post-service/pkg/logger"
	grpcClient "post-service/service/grpc_client"
	"post-service/storage"
	"strconv"

	"github.com/jmoiron/sqlx"
	// "go.mongodb.org/mongo-driver/mongo"
)

// PostService ...
type PostService struct {
	client  grpcClient.IServiceManager
	storage storage.IStorage
	logger  l.Logger
}

// NewPostService ... Postgres
func NewPostService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// NewPostService ... Mongo
// func NewPostService(db *mongo.Client, log l.Logger) *PostService {
// 	return &PostService{
// 		storage: storage.NewStoragePg(db),
// 		logger:  log,
// 	}
// }

// CategoryIcon Start

func (s *PostService) CreateCategoryIcon(ctx context.Context, req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	res, err := s.storage.Post().CreateCategoryIconDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ReadCategoryIcon(ctx context.Context, req *pbp.IdRequest) (*pbp.CategoryIcon, error) {
	res, err := s.storage.Post().ReadCategoryIconDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) UpdateCategoryIcon(ctx context.Context, req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	res, err := s.storage.Post().UpdateCategoryIconDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) DeleteCategoryIcon(ctx context.Context, req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	res, err := s.storage.Post().DeleteCategoryIconDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ListCategoryIcons(ctx context.Context, req *pbp.GetAllRequest) (*pbp.ListCategoryIconResponse, error) {
	CategoryIcons, err := s.storage.Post().ListCategoryIconsDB(req)
	if err != nil {
		return nil, err
	}
	return CategoryIcons, nil
}

// CategoryIcon End

// Category Start
func (s *PostService) CreateCategory(ctx context.Context, req *pbp.Category) (*pbp.Category, error) {
	res, err := s.storage.Post().CreateCategoryDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ReadCategory(ctx context.Context, req *pbp.IdRequest) (*pbp.Category, error) {
	res, err := s.storage.Post().ReadCategoryDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) UpdateCategory(ctx context.Context, req *pbp.Category) (*pbp.Category, error) {
	res, err := s.storage.Post().UpdateCategoryDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) DeleteCategory(ctx context.Context, req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	res, err := s.storage.Post().DeleteCategoryDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ListCategorys(ctx context.Context, req *pbp.GetAllRequest) (*pbp.ListCategoryResponse, error) {
	Categorys, err := s.storage.Post().ListCategorysDB(req)
	if err != nil {
		return nil, err
	}
	return Categorys, nil
}

// Category End

// Product Start
func (s *PostService) CreateProduct(ctx context.Context, req *pbp.Product) (*pbp.Product, error) {
	res, err := s.storage.Post().CreateProductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ReadProduct(ctx context.Context, req *pbp.IdRequest) (*pbp.Product, error) {
	res, err := s.storage.Post().ReadProductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) UpdateProduct(ctx context.Context, req *pbp.Product) (*pbp.Product, error) {
	res, err := s.storage.Post().UpdateProductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) DeleteProduct(ctx context.Context, req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	res, err := s.storage.Post().DeleteProductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ListProducts(ctx context.Context, req *pbp.GetAllRequest) (*pbp.ListProductResponse, error) {
	products, err := s.storage.Post().ListProductsDB(req)
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s *PostService) ListProductsWithComments(ctx context.Context, req *pbp.GetAllRequest) (*pbp.ListProductWithCommentResponse, error) {
	// var productsWithComments pbp.ListProductWithCommentResponse
	products, err := s.storage.Post().ListProductsWithCommentsDB(req)
	if err != nil {
		return nil, err
	}

	for v, k := range products.Productwithcomments {
		strid := strconv.Itoa(int(k.Id))

		comments, err := s.client.CommentService().ListCommentsByProductId(ctx, &pbc.IdRequest{Id: strid})
		if err != nil {
			return nil, err
		}
		for _, j := range comments.Comments {
			productComment := pbp.Comment{
				Id:        j.Id,
				Content:   j.Content,
				UserId:    j.UserId,
				ProductId: j.ProductId,
				CreatedAt: j.CreatedAt,
				UpdatedAt: j.UpdatedAt,
			}
			products.Productwithcomments[v].Comments = append(products.Productwithcomments[v].Comments, &productComment)
		}
	}
	return products, nil
}

// Product End

// Orderproduct Start

func (s *PostService) CreateOrderproduct(ctx context.Context, req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
	res, err := s.storage.Post().CreateOrderproductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ReadOrderproduct(ctx context.Context, req *pbp.IdRequest) (*pbp.Orderproduct, error) {
	res, err := s.storage.Post().ReadOrderproductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) UpdateOrderproduct(ctx context.Context, req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
	res, err := s.storage.Post().UpdateOrderproductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) DeleteOrderproduct(ctx context.Context, req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	res, err := s.storage.Post().DeleteOrderproductDB(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PostService) ListOrderproducts(ctx context.Context, req *pbp.GetAllRequest) (*pbp.ListOrderproductResponse, error) {
	Orderproducts, err := s.storage.Post().ListOrderproductsDB(req)
	if err != nil {
		return nil, err
	}
	return Orderproducts, nil
}

// Orderproduct End
