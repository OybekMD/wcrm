package main

import (
	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/kafka"
	"api-gateway/pkg/logger"
	k "api-gateway/pkg/kafka"
	"api-gateway/services"
	"fmt"

	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	// With csv file code
	fileAdapter := fileadapter.NewAdapter("./config/auth.csv")

	enforcer, err := casbin.NewEnforcer("./config/auth.conf", fileAdapter)
	if err != nil {
		log.Error("NewEnforcer error", logger.Error(err))
		return
	}

	kafka.Init(k.NewProducerInit(cfg, log))

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		Enforcer:       enforcer,
		ServiceManager: serviceManager,
	})

	fmt.Println("\n\x1b[34mWRCM Server Run!\x1b[0m\n")
	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}

	// Graceful Shuttingdown
	// go func() {
	// 	if err := server.Run(cfg.HTTPPort); err != nil {
	// 		log.Fatal("failed to run http server", logger.Error(err))
	// 		panic(err)
	// 	}
	// }()

	// fmt.Println("\x1b[32mWRCM Started\x1b[0m")

	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	// <-quit

	// fmt.Println("\x1b[32mWRCM Graceful Shutting Down\x1b[0m")
}
