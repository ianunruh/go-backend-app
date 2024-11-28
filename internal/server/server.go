package server

import (
	"context"
	"errors"
	"net/http"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"github.com/ianunruh/go-backend-app/internal/generated/api"
	"github.com/ianunruh/go-backend-app/internal/httpapi"
)

func Run(
	ctx context.Context,
	cfg Config,
	meterProvider *sdkmetric.MeterProvider,
	tracerProvider *sdktrace.TracerProvider,
	log *zap.Logger,
) error {
	handlers := httpapi.NewHandlers()

	srv, err := api.NewServer(handlers,
		api.WithMeterProvider(meterProvider),
		api.WithTracerProvider(tracerProvider))
	if err != nil {
		return err
	}

	log.Info("Starting API server", zap.String("addr", cfg.ListenAddr))

	// TODO use context
	if err := http.ListenAndServe(cfg.ListenAddr, srv); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
