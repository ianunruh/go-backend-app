package server

import (
	"github.com/ianunruh/go-backend-app/internal/server/requestlog"
)

type Config struct {
	ListenAddr string `yaml:"listenAddr" env:"LISTEN_ADDR" envDefault:"localhost:9080"`

	RequestLog requestlog.Config `yaml:"requestLog" env:"REQUEST_LOG"`
}
