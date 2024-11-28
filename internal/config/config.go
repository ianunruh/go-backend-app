package config

import (
	"os"

	"github.com/caarlos0/env/v8"
	"gopkg.in/yaml.v3"

	"github.com/ianunruh/go-backend-app/internal/debug"
	"github.com/ianunruh/go-backend-app/internal/server"
	"github.com/ianunruh/go-backend-app/internal/telemetry"
)

type Config struct {
	Debug debug.Config `yaml:"debug" envPrefix:"DEBUG_"`

	Log telemetry.LogConfig `yaml:"log" envPrefix:"LOG_"`

	Metrics telemetry.MetricsConfig `yaml:"metrics" envPrefix:"METRICS_"`

	Server server.Config `yaml:"server" envPrefix:"SERVER_"`

	Tracing telemetry.TracingConfig `yaml:"tracing" envPrefix:"TRACING_"`
}

func Load(path string) (*Config, error) {
	var cfg Config

	envOpts := env.Options{
		Prefix: "APP_",
	}
	if err := env.ParseWithOptions(&cfg, envOpts); err != nil {
		return nil, err
	}

	if err := LoadFile(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadFile(path string, out *Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	if err := decoder.Decode(out); err != nil {
		return err
	}

	return nil
}
