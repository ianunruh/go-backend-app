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
	"github.com/ianunruh/go-backend-app/internal/server/requestlog"
)

func Run(
	ctx context.Context,
	cfg Config,
	meterProvider *sdkmetric.MeterProvider,
	tracerProvider *sdktrace.TracerProvider,
	log *zap.Logger,
) error {
	handlers := httpapi.NewHandlers()

	apiSrv, err := api.NewServer(handlers,
		api.WithMeterProvider(meterProvider),
		api.WithTracerProvider(tracerProvider))
	if err != nil {
		return err
	}

	h := requestlog.Middleware(apiSrv, cfg.RequestLog, log)
	h = tracePropagationMiddleware(h)

	log.Info("Starting API server", zap.String("addr", cfg.ListenAddr))

	// TODO use context
	if err := http.ListenAndServe(cfg.ListenAddr, h); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
	return nil
}
