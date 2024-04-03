package mongo

import (
	pbc "comment-service/genproto/comment"
	"comment-service/pkg/logger"
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentMongoRepo struct {
	db *mongo.Client
}

// NewUserRepo ...
func NewCommentMongoRepo(db *mongo.Client) *commentMongoRepo {
	return &commentMongoRepo{db: db}
}

type Comment struct {
	Id        int64
	Content   string
	UserId    string
	ProductId int64
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

func (r *commentMongoRepo) CreateCommentDB(user *pbc.Comment) (*pbc.Comment, error) {
	nowTime := time.Now()
	userWrite := &Comment{
		Id:        1,
		Content:   user.Content,
		UserId:    user.UserId,
		ProductId: user.ProductId,
		CreatedAt: nowTime.String(),
		UpdatedAt: nowTime.String(),
	}
	_, err := r.db.Database("wcrmdb").Collection("comments").InsertOne(context.Background(), userWrite)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	user.Id = userWrite.Id
	user.CreatedAt = userWrite.CreatedAt
	user.UpdatedAt = userWrite.UpdatedAt

	return user, nil
}

func (r *commentMongoRepo) ReadCommentDB(req *pbc.IdRequest) (*pbc.Comment, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)
	
	collection := r.db.Database("wcrmdb").Collection("comments")

	result := pbc.Comment{}
	filter := bson.M{"id": id64}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *commentMongoRepo) UpdateCommentDB(res *pbc.Comment) (*pbc.Comment, error) {
	var resp pbc.Comment

	filter := bson.D{{Key: "id", Value: res.Id}}

	nowTime := time.Now().Format(time.RFC3339) // Format timestamp correctly
	update := bson.D{
		{
			Key: "$set", Value: bson.D{
				{Key: "content", Value: res.Content},
				{Key: "user_id", Value: res.UserId},
				{Key: "product_id", Value: res.ProductId},
				{Key: "updated_at", Value: nowTime},
			},
		},
	}

	// Perform FindOneAndUpdate operation
	result := r.db.Database("wcrmdb").Collection("comments").FindOneAndUpdate(context.TODO(), filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			logger.Error(result.Err())
			return nil, result.Err()
		}
		logger.Error(result.Err())
		return nil, result.Err()
	}

	// Decode the updated document
	if err := result.Decode(&resp); err != nil {
		logger.Error(err)
		return nil, err
	}

	return &resp, nil
}

func (r *commentMongoRepo) DeleteCommentDB(req *pbc.IdRequest) (*pbc.MessageResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, err
	}
	id64 := int64(id)

	_, err = r.db.Database("wcrmdb").Collection("comments").DeleteOne(context.TODO(), bson.D{{Key: "id", Value: id64}})
	if err != nil {
		logger.Error(err)
		return &pbc.MessageResponse{Message: "Comment not found for deletion."}, err
	}

	return &pbc.MessageResponse{Message: "Comment successfully deleted!"}, nil
}

func (r *commentMongoRepo) ListCommentsDB(req *pbc.GetAllRequest) (*pbc.ListCommentResponse, error) {
	skip := (req.Page - 1) * req.Limit

	cursor, err := r.db.Database("wcrmdb").Collection("comments").Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: &req.Limit,
		Skip:  &skip,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var allComments pbc.ListCommentResponse
	for cursor.Next(context.TODO()) {
		var result pbc.Comment
		err := cursor.Decode(&result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allComments.Comments = append(allComments.Comments, &result)
	}

	return &allComments, nil
}

func (r *commentMongoRepo) ListCommentsByProductIdDB(req *pbc.IdRequest) (*pbc.ListCommentResponse, error) {
	var allComments pbc.ListCommentResponse

	collection := r.db.Database("your_database_name").Collection("comments")

	filter := bson.M{
		"product_id": req.Id,
		"deleted_at": nil,
	}

	options := options.Find()

	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}
		allComments.Comments = append(allComments.Comments, &comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &allComments, nil
}
