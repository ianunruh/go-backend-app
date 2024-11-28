package telemetry

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

type LogConfig struct {
	Dev bool `yaml:"dev" env:"DEV"`

	Level string `yaml:"level" env:"LEVEL" envDefault:"info"`

	DumpConfig bool `yaml:"dumpConfig" env:"DUMP_CONFIG"`
}

func NewLog(cfg LogConfig) (*zap.Logger, zap.AtomicLevel, error) {
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		return nil, zap.AtomicLevel{}, err
	}

	zapCfg := newZapConfig(cfg)
	zapCfg.Level = level

	log, err := zapCfg.Build()
	if err != nil {
		return nil, zap.AtomicLevel{}, err
	}

	return log, level, nil
}

func newZapConfig(cfg LogConfig) zap.Config {
	if cfg.Dev {
		return zap.NewDevelopmentConfig()
	}
	return zap.NewProductionConfig()
}

func SyncLog(log *zap.Logger) {
	if log == nil {
		return
	}

	if err := log.Sync(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to sync log:", err)
	}
}
