package main

import (
	"fmt"
	"net"
	"user-service/config"
	pbu "user-service/genproto/user"
	"user-service/pkg/db"
	"user-service/pkg/logger"
	"user-service/service"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	
	log := logger.New(cfg.LogLevel, "user-service")
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

	userService := service.NewUserService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pbu.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

func consumerHandler(message []byte) {
	fmt.Println(string(message))
}
