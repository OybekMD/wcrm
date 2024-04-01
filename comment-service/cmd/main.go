package main

import (
	"google.golang.org/grpc"
	"net"
	"comment-service/config"
	pbc "comment-service/genproto/comment"
	"comment-service/pkg/db"
	"comment-service/pkg/logger"
	"comment-service/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	// Postgres
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	// Mongo
	// connDB, err := db.ConnectToMongoDB(cfg)
	// if err != nil {
	// 	log.Fatal("sqlx connection to postgres error", logger.Error(err))
	// }

	postService := service.NewCommentService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbc.RegisterCommentServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
