package mongo

import (
	"context"
	pbp "post-service/genproto/post"
	"post-service/pkg/logger"
	"strconv"

	// "strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Category struct {
	Id        int64
	Name      string
	IconId    int64
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (r *postMongoRepo) CreateCategoryDB(req *pbp.Category) (*pbp.Category, error) {
	collection := r.db.Database("wcrmdb").Collection("categories")
	now := time.Now().String()
	creq := Category{
		Id:        1,
		Name:      req.Name,
		IconId:    1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err := collection.InsertOne(context.Background(), creq)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	req.Id = creq.Id
	req.IconId = creq.IconId
	req.CreatedAt = creq.CreatedAt
	req.UpdatedAt = creq.UpdatedAt

	return req, nil
}

func (r *postMongoRepo) ReadCategoryDB(req *pbp.IdRequest) (*pbp.Category, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)
	
	collection := r.db.Database("wcrmdb").Collection("categories")

	result := pbp.Category{}
	filter := bson.M{"id": id64, "deleted_at": nil}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	

	return &result, nil
}

func (r *postMongoRepo) UpdateCategoryDB(req *pbp.Category) (*pbp.Category, error) {
	collection := r.db.Database("wcrmdb").Collection("categories")

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

func (r *postMongoRepo) DeleteCategoryDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	collection := r.db.Database("wcrmdb").Collection("categories")

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
		return &pbp.MessageResponse{Message: "Category not found for deletion."}, nil
	}

	return &pbp.MessageResponse{Message: "Category successfully deleted!"}, nil
}

func (r *postMongoRepo) ListCategorysDB(req *pbp.GetAllRequest) (*pbp.ListCategoryResponse, error) {
	skip := (req.Page - 1) * req.Limit
	collection := r.db.Database("wcrmdb").Collection("categories")

	var allCategories pbp.ListCategoryResponse
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
		var category pbp.Category
		err := cursor.Decode(&category)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allCategories.Categorys = append(allCategories.Categorys, &category)
	}

	return &allCategories, nil
}
