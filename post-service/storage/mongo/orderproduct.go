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

type Orderproduct struct {
	Id        int64
	UserId    string
	ProductId int64
	CreatedAt string
}

func (r *postMongoRepo) CreateOrderproductDB(req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
	collection := r.db.Database("wcrmdb").Collection("orderproducts")
	now := time.Now().String()
	creq := Orderproduct{
		Id:        1,
		UserId:    req.UserId,
		ProductId: req.ProductId,
		CreatedAt: now,
	}

	_, err := collection.InsertOne(context.Background(), creq)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	req.Id = creq.Id
	req.CreatedAt = creq.CreatedAt

	return req, nil
}

func (r *postMongoRepo) ReadOrderproductDB(req *pbp.IdRequest) (*pbp.Orderproduct, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)

	collection := r.db.Database("wcrmdb").Collection("orderproducts")

	var result pbp.Orderproduct
	filter := bson.M{"id": id64}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *postMongoRepo) UpdateOrderproductDB(req *pbp.Orderproduct) (*pbp.Orderproduct, error) {
	collection := r.db.Database("wcrmdb").Collection("orderproducts")

	filter := bson.M{"_id": req.Id}
	update := bson.M{"$set": req}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return req, nil
}

func (r *postMongoRepo) DeleteOrderproductDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	collection := r.db.Database("wcrmdb").Collection("orderproducts")

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
		return &pbp.MessageResponse{Message: "CategoryIcon not found for deletion."}, nil
	}

	return &pbp.MessageResponse{Message: "Orderproduct successfully deleted!"}, nil
}

func (r *postMongoRepo) ListOrderproductsDB(req *pbp.GetAllRequest) (*pbp.ListOrderproductResponse, error) {
	skip := (req.Page - 1) * req.Limit
	collection := r.db.Database("wcrmdb").Collection("orderproducts")

	var allOrderProducts pbp.ListOrderproductResponse
	cursor, err := collection.Find(context.Background(), bson.M{}, &options.FindOptions{
		Limit: &req.Limit,
		Skip:  &skip,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var orderProduct pbp.Orderproduct
		err := cursor.Decode(&orderProduct)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allOrderProducts.Orderproducts = append(allOrderProducts.Orderproducts, &orderProduct)
	}

	return &allOrderProducts, nil
}
