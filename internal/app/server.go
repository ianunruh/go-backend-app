package app

import (
	"context"

	"github.com/ianunruh/go-backend-app/internal/server"
)

func (ct *Container) RunServer(ctx context.Context) error {
	return server.Run(ctx, ct.Cfg.Server, ct.MeterProvider, ct.TracerProvider, ct.Log)
}
