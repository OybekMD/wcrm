package storage

import (
	"user-service/storage/postgres"
	"user-service/storage/repo"

	
	"github.com/jmoiron/sqlx"
	// mongoDB "user-service/storage/mongo"
	// "go.mongodb.org/mongo-driver/mongo"
)

// IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

// Swap from postgres to mongo
type storagePg struct {
	// db       *mongo.Client
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

//Postger
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

// Mongo
// func NewStoragePg(db *mongo.Client) *storagePg {
// 	return &storagePg{
// 		db:       db,
// 		userRepo: mongoDB.NewUserMongoRepo(db),
// 	}
// }

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
