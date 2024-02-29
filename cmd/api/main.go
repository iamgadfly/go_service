package main

import (
	"go.uber.org/zap"
	"go_service/config"
	_ "go_service/docs"
	"go_service/internal/grpcserver"
	"go_service/internal/server"
	"go_service/pkg/psql"
	"go_service/pkg/redis"
	"log"
)

// @title App GRPC and REST
// @version 1.0
// @description app
// @host localhost:8000
// @BasePath /
func main() {
	cfgFile, err := config.LoadConfig("config/config-local.yml")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	db, err := psql.NewPsqlDB(cfg)
	if err != nil {
		log.Fatalf("psql connect error: %v\n", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("zapp logger error: %v\n", err)
	}
	sugar := logger.Sugar()

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	sugar.Info("Redis connected")

	er := make(chan error)

	grpcServer := grpcserver.NewServerGRPC(db, redisClient, cfg, sugar)
	s := server.NewServer(cfg, redisClient, cfg.Server.Port, db, sugar)

	go grpcServer.Run(er)
	go s.Run(er)

	select {
	case err = <-er:
		if err != nil {
			log.Fatalf("Server error: %v\n", err.Error())
		}
	}
}
