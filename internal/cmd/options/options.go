package options

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/ianunruh/go-backend-app/internal/app"
	"github.com/ianunruh/go-backend-app/internal/config"
	"github.com/ianunruh/go-backend-app/internal/telemetry"
)

type Options struct {
	ConfigPath string

	LogLevel string
}

func (opts *Options) NewContainer() (*app.Container, error) {
	cfg, err := opts.newConfig()
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}

	log, level, err := opts.newLog(cfg)
	if err != nil {
		return nil, fmt.Errorf("creating log: %w", err)
	}

	if cfg.Log.DumpConfig {
		log.Debug("Config loaded", zap.Any("config", cfg))
	}

	return app.NewContainer(cfg, log, level)
}

func (opts *Options) newConfig() (*config.Config, error) {
	path := os.Getenv("APP_CONFIG_PATH")
	if opts.ConfigPath != "" {
		path = opts.ConfigPath
	}
	if path == "" {
		path = "config.yaml"
	}

	return config.Load(path)
}

func (opts *Options) newLog(cfg *config.Config) (*zap.Logger, zap.AtomicLevel, error) {
	if opts.LogLevel != "" {
		cfg.Log.Level = opts.LogLevel
	}

	return telemetry.NewLog(cfg.Log)
}
