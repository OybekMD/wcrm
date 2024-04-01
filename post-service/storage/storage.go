package storage

import (
	"github.com/jmoiron/sqlx"
	"post-service/storage/postgres"
	"post-service/storage/repo"
	// "go.mongodb.org/mongo-driver/mongo"
	// mongoDB "post-service/storage/mongo"
)

// IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

// Swap postgres to mongo
type storagePg struct {
	// db       *mongo.Client
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

// NewStoragePg ...

// Postgres
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
	}
}

// Mongo
// func NewStoragePg(db *mongo.Client) *storagePg {
// 	return &storagePg{
// 		db:       db,
// 		postRepo: mongoDB.NewPostMongoRepo(db),
// 	}
// }

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
