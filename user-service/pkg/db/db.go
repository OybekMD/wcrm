package db

import (
	"context"
	"fmt"
	"log"
	"user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}

	return connDb, nil
}

func ConnectToRedisDB(cfg config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb, nil
}

func ConnectToMongoDB(cfg config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return client, nil
}

func ConnectDBForSuite(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, nil
	}

	cleanUpFunc := func() {
		connDb.Close()
	}

	return connDb, cleanUpFunc
}
