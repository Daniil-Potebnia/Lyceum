package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"service/internal/config"
	"service/internal/repository"
	service "service/internal/service"
	"service/internal/transport/grpc"
	"service/pkg/db/cache"
	"service/pkg/db/postgres"
	"service/pkg/logger"
	"syscall"
)

const (
	serviceName = "task"
)

func main() {
	ctx := context.Background()
	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)
	cfg := config.New()
	if cfg == nil {
		panic("load_config")
	}

	db, err := postgres.New(cfg.Config)
	if err != nil {
		panic(err)
	}

	redis := cache.New(cfg.RedisConfig)
	fmt.Println(redis.Ping(ctx))

	repo := repository.NewOrderRepository(db)

	srv := service.NewOrderService(repo)
	
	grpcserver, err := grpc.New(ctx, cfg.GRPCServerPort, cfg.RestServerPort,  srv)
	if err != nil {
		mainLogger.Error(ctx, err.Error())
		return
	}
	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := grpcserver.Start(ctx); err != nil {
			mainLogger.Error(ctx, err.Error())
		}
	}()

	<-graceCh

	if err := grpcserver.Stop(ctx); err != nil {
		mainLogger.Error(ctx, err.Error())
	}
	mainLogger.Info(ctx, "Server stopped")
	fmt.Println("Server Stopped")
}
