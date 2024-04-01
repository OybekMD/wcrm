package mongo

import (
	"context"
	"time"
	pbu "user-service/genproto/user"
	"user-service/pkg/logger"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userMongoRepo struct {
	db *mongo.Client
}

// NewUserRepo ...
func NewUserMongoRepo(db *mongo.Client) *userMongoRepo {
	return &userMongoRepo{db: db}
}

func (r *userMongoRepo) CreateUserDB(user *pbu.User) (*pbu.User, error) {
	if user.Id == "" {
		user.Id = uuid.New().String()
	}

	nowTime := time.Now()
	userWrite := &pbu.User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		PhoneNumber:  user.PhoneNumber,
		Bio:          user.Bio,
		BirthDay:     user.BirthDay,
		Email:        user.Email,
		Avatar:       user.Avatar,
		Password:     user.Password,
		RefreshToken: user.RefreshToken,
		CreatedAt:    nowTime.String(),
		UpdatedAt:    nowTime.String(),
	}
	
	_, err := r.db.Database("wcrmdb").Collection("user_info").InsertOne(context.Background(), userWrite)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return userWrite, nil
}

func (r *userMongoRepo) ReadUserDB(req *pbu.IdRequest) (*pbu.User, error) {

	filter := bson.D{{Key: "id", Value: req.Id}}
	var result pbu.User
	err := r.db.Database("wcrmdb").Collection("user_info").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return &result, nil
}

func (r *userMongoRepo) UpdateUserDB(res *pbu.User) (*pbu.User, error) {
    var resp pbu.User

    filter := bson.D{{Key: "id", Value: res.Id}}

    nowTime := time.Now().Format(time.RFC3339) // Format timestamp correctly
    update := bson.D{
        {
            Key: "$set", Value: bson.D{
                {Key: "first_name", Value: res.FirstName},
                {Key: "last_name", Value: res.LastName},
                {Key: "username", Value: res.Username},
                {Key: "bio", Value: res.Bio},
                {Key: "birth_day", Value: res.BirthDay},
                {Key: "email", Value: res.Email},
                {Key: "avatar", Value: res.Avatar},
                {Key: "password", Value: res.Password},
                {Key: "refresh_token", Value: res.RefreshToken},
                {Key: "updated_at", Value: nowTime},
            },
        },
    }

    // Perform FindOneAndUpdate operation
    result := r.db.Database("wcrmdb").Collection("user_info").FindOneAndUpdate(context.TODO(), filter, update)
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


func (r *userMongoRepo) DeleteUserDB(req *pbu.IdRequest) (*pbu.MessageResponse, error) {
	_, err := r.db.Database("wcrmdb").Collection("user_info").DeleteOne(context.TODO(), bson.D{{Key: "id", Value: req.Id}})
	if err != nil {
		logger.Error(err)
		return &pbu.MessageResponse{Message: "User not found for deletion."}, err
	}
	
	return &pbu.MessageResponse{Message: "User successfully deleted!"}, nil
}

func (r *userMongoRepo) ListUserDB(req *pbu.GetAllRequest) (*pbu.ListUserResponse, error) {
	skip := (req.Page - 1) * req.Limit

	cursor, err := r.db.Database("wcrmdb").Collection("user_info").Find(context.TODO(), bson.M{}, &options.FindOptions{
		Limit: &req.Limit,
		Skip:  &skip,
	})
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var allUsers pbu.ListUserResponse
	for cursor.Next(context.TODO()) {
		var result pbu.User
		err := cursor.Decode(&result)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		allUsers.Users = append(allUsers.Users, &result)
	}

	return &allUsers, nil
}

func (r *userMongoRepo) CheckUniqueDB(req *pbu.CheckUniqueRequest) (*pbu.CheckUniqueRespons, error) {
	var response pbu.User
	filter := bson.M{req.Column: req.Value}
	err := r.db.Database("wcrmdb").Collection("user_info").FindOne(context.Background(), filter).Decode(&response)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return &pbu.CheckUniqueRespons{
				IsExist: false,
			}, nil
		}
		return nil, err
	}

	if response.Id == "" {
		return &pbu.CheckUniqueRespons{
			IsExist: false,
		}, nil
	}

	return &pbu.CheckUniqueRespons{
		IsExist: true,
	}, nil
}
