package work

import (
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

type Config struct {
	Concurrency int `env:"CONCURRENCY" envDefault:"10"`
}

func AsynqConfig(cfg Config, log *zap.Logger) asynq.Config {
	return asynq.Config{
		Concurrency: cfg.Concurrency,
		Logger:      newAsynqLogger(log),
		Queues: map[string]int{
			"critical": 4,
			"high":     3,
			"default":  2,
			"low":      1,
		},
	}
}
