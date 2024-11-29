package config

import (
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Address  string `yaml:"address" env:"ADDRESS"`
	Password string `yaml:"password" env:"PASSWORD"`
	DB       int    `yaml:"db" env:"DB"`
}

func AsynqRedisClientOpt(cfg Redis) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	}
}

func RedisClient(cfg Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
