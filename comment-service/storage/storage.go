package storage

import (
	"comment-service/storage/postgres"
	"comment-service/storage/repo"

	"github.com/jmoiron/sqlx"
	// "go.mongodb.org/mongo-driver/mongo"
	// mongoDB "comment-service/storage/mongo"
)

// IStorage ...
type IStorage interface {
	Comment() repo.CommentStorageI
}

// Swap from postgres to mongo
type storagePg struct {
	// db       *mongo.Client
	db          *sqlx.DB
	commentRepo repo.CommentStorageI
}

// NewStoragePg ...

// Postgres
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

// Mongo
// func NewStoragePg(db *mongo.Client) *storagePg {
// 	return &storagePg{
// 		db:       db,
// 		commentRepo: mongoDB.NewCommentMongoRepo(db),
// 	}
// }

func (s storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}
