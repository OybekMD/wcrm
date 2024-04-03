package main

import (
	"net"
	"post-service/config"
	pbc "post-service/genproto/post"
	"post-service/pkg/db"
	"post-service/pkg/logger"
	"post-service/service"
	grpcclient "post-service/service/grpc_client"
	

	"google.golang.org/grpc"
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
	grpcClient, err := grpcclient.New(cfg)
	if err != nil {
		log.Fatal("grpc client dail error", logger.Error(err))
	}

	postService := service.NewPostService(connDB, log, grpcClient)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbc.RegisterPostServiceServer(s, postService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
