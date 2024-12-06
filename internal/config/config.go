package config

import (
	"os"
	"service/pkg/db/cache"
	"service/pkg/db/postgres"
	"strconv"
)

type Config struct {
	postgres.Config
	cache.RedisConfig

	GRPCServerPort int `env:"GRPC_SERVER_PORT" env-default:"50051"`
	RestServerPort int `env:"REST_SERVER_PORT" env-default:"50050"`
}

func New() *Config {
	cfg := Config{}

	grpc, err := strconv.Atoi(os.Getenv("GRPC_SERVER_PORT"))
	if err != nil {
		return nil
	}
	cfg.GRPCServerPort = grpc
	rest, err := strconv.Atoi(os.Getenv("REST_SERVER_PORT"))
	if err != nil {
		return nil
	}
	cfg.RestServerPort = rest

	cfg.Config.DBName = os.Getenv("POSTGRES_DB")
	cfg.Config.Host = os.Getenv("POSTGRES_HOST")
	cfg.Config.Port = os.Getenv("POSTGRES_PORT")
	cfg.Config.UserName = os.Getenv("POSTGRES_USER")
	cfg.Config.Password = os.Getenv("POSTGRES_PASSWORD")

	cfg.RedisConfig.Host = os.Getenv("REDIS_HOST")
	cfg.RedisConfig.Port = os.Getenv("REDIS_PORT")

	return &cfg
}
