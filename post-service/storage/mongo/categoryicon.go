package mongo

import (
	"context"
	pbp "post-service/genproto/post"
	"post-service/pkg/logger"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type postMongoRepo struct {
	db *mongo.Client
}

// NewPostRepo creates a new instance of postMongoRepo.
func NewPostMongoRepo(db *mongo.Client) *postMongoRepo {
	return &postMongoRepo{db: db}
}

type CategoryIcon struct {
	Id      int64
	Name    string
	Picture string
}

func (r *postMongoRepo) CreateCategoryIconDB(req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	collection := r.db.Database("wcrmdb").Collection("category_icons")
	creq := CategoryIcon{
		Id:      1,
		Name:    req.Name,
		Picture: req.Picture,
	}

	_, err := collection.InsertOne(context.Background(), creq)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	req.Id = creq.Id

	return req, nil
}

func (r *postMongoRepo) ReadCategoryIconDB(req *pbp.IdRequest) (*pbp.CategoryIcon, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)

	collection := r.db.Database("wcrmdb").Collection("category_icons")

	result := pbp.CategoryIcon{}
	filter := bson.M{"id": id64}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &result, nil
	// return &pbp.CategoryIcon{Id: result.Id, Name: result.Name, Picture: result.Picture}, nil
}

func (r *postMongoRepo) UpdateCategoryIconDB(req *pbp.CategoryIcon) (*pbp.CategoryIcon, error) {
	collection := r.db.Database("wcrmdb").Collection("category_icons")
	filter := bson.M{"_id": req.Id}
	update := bson.M{"$set": req}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return req, nil
}

func (r *postMongoRepo) DeleteCategoryIconDB(req *pbp.IdRequest) (*pbp.MessageResponse, error) {
	collection := r.db.Database("wcrmdb").Collection("category_icons")

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

	return &pbp.MessageResponse{Message: "CategoryIcon successfully deleted!"}, nil
}

func (r *postMongoRepo) ListCategoryIconsDB(req *pbp.GetAllRequest) (*pbp.ListCategoryIconResponse, error) {
	skip := (req.Page - 1) * req.Limit

	collection := r.db.Database("wcrmdb").Collection("category_icons")

	var allCategoryIcons pbp.ListCategoryIconResponse
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
		var CategoryIcon pbp.CategoryIcon
		err := cursor.Decode(&CategoryIcon)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allCategoryIcons.Categoryicons = append(allCategoryIcons.Categoryicons, &CategoryIcon)
	}

	return &allCategoryIcons, nil
}
