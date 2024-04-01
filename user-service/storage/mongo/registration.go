package mongo

import (
	"context"
	"errors"
	"log"
	"time"
	pbu "user-service/genproto/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *userMongoRepo) LoginDB(req *pbu.LoginRequest) (*pbu.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user pbu.User

	collection := r.db.Database("wcrmdb").Collection("user_info")

	filter := bson.M{
		"$or": []bson.M{
			{"email": req.Email},
			{"username": req.Email},
		},
		"deleted_at": nil,
	}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userMongoRepo) UpdatePasswordDB(req *pbu.UpdatePasswordRequest) (*pbu.MessageResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := r.db.Database("your_database_name").Collection("your_collection_name")

	filter := bson.M{"email": req.Email}
	update := bson.M{"$set": bson.M{"password_hash": req.Password}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("error while updating user password:", err)
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return nil, errors.New("user not found")
	}

	res := &pbu.MessageResponse{
		Message: "Password Successfully updated!",
	}
	return res, nil
}


func (r *userMongoRepo) GetFullNameDB(req *pbu.LoginRequest) (*pbu.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user pbu.User

	collection := r.db.Database("wcrmdb").Collection("user_info")

	filter := bson.M{
		"email":      req.Email,
		"deleted_at": nil,
	}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}