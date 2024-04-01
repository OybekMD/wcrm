package mongo

import (
	"context"
	pbp "post-service/genproto/post"
	"post-service/pkg/logger"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Product struct {
	Id          int64
	Title       string
	Description string
	Price       int64
	Picture     string
	CategoryId  int64
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

func (r *postMongoRepo) CreateProductDB(req *pbp.Product) (*pbp.Product, error) {
	collection := r.db.Database("wcrmdb").Collection("products")
	now := time.Now().String()

	creq := Product{
		Id: 1,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		Picture: req.Picture,
		CategoryId: req.CategoryId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := collection.InsertOne(context.Background(), creq)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	req.Id = creq.Id
	req.CreatedAt = creq.CreatedAt
	req.UpdatedAt = creq.UpdatedAt

	return req, nil
}

func (r *postMongoRepo) ReadProductDB(req *pbp.IdRequest) (*pbp.Product, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)

	collection := r.db.Database("wcrmdb").Collection("products")

	var result pbp.Product
	filter := bson.M{"id": id64, "deleted_at": nil}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *postMongoRepo) UpdateProductDB(req *pbp.Product) (*pbp.Product, error) {
	collection := r.db.Database("wcrmdb").Collection("products")

	now := time.Now().String()
	req.UpdatedAt = now

	filter := bson.M{"_id": req.Id}
	update := bson.M{"$set": req}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return req, nil
}

func (r *postMongoRepo) DeleteProductDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	collection := r.db.Database("wcrmdb").Collection("products")

	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)

	filter := bson.M{"id": id64}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if result.DeletedCount == 0 {
		return &pbp.MessageResponse{Message: "Product not found for deletion."}, nil
	}

	return &pbp.MessageResponse{Message: "Product successfully deleted!"}, nil
}

func (r *postMongoRepo) ListProductsDB(req *pbp.GetAllRequest) (*pbp.ListProductResponse, error) {
	skip := (req.Page - 1) * req.Limit
	collection := r.db.Database("wcrmdb").Collection("products")

	var allProducts pbp.ListProductResponse
	filter := bson.M{"deleted_at": nil}
	cursor, err := collection.Find(context.Background(), filter, &options.FindOptions{
		Limit: &req.Limit,
		Skip:  &skip,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product pbp.Product
		err := cursor.Decode(&product)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allProducts.Products = append(allProducts.Products, &product)
	}

	return &allProducts, nil
}
