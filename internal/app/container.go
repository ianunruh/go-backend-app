package app

import (
	"context"
	"fmt"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"

	"github.com/hibiken/asynq"
	"github.com/ianunruh/go-backend-app/internal/config"
	"github.com/ianunruh/go-backend-app/internal/debug"
	"github.com/ianunruh/go-backend-app/internal/telemetry"
	"github.com/ianunruh/go-backend-app/internal/work"
)

func NewContainer(cfg *config.Config, log *zap.Logger, logLevel zap.AtomicLevel) (*Container, error) {
	debugServer := debug.NewServer(cfg.Debug, log, logLevel)
	if err := debugServer.Start(); err != nil {
		return nil, fmt.Errorf("starting debug server: %w", err)
	}

	meterProvider, err := telemetry.NewOTELMeterProvider()
	if err != nil {
		return nil, fmt.Errorf("creating OTEL meter provider: %w", err)
	}

	traceExporter, err := telemetry.NewOTLPTraceExporter(context.Background(), cfg.Tracing)
	if err != nil {
		return nil, fmt.Errorf("creating OTLP exporter: %w", err)
	}

	tracerProvider, err := telemetry.NewOTELTracerProvider(context.Background(), traceExporter)
	if err != nil {
		return nil, fmt.Errorf("creating OTEL trace provider: %w", err)
	}

	metricsServer := telemetry.NewMetricsServer(cfg.Metrics, log)
	if err := metricsServer.Start(); err != nil {
		return nil, fmt.Errorf("starting metrics server: %w", err)
	}

	redisOpt := config.AsynqRedisClientOpt(cfg.Redis)

	workQueue := work.NewQueue(redisOpt, meterProvider, tracerProvider, log)

	ct := &Container{
		Cfg:      cfg,
		Log:      log,
		LogLevel: logLevel,

		MeterProvider:  meterProvider,
		TracerProvider: tracerProvider,

		DebugServer:   debugServer,
		MetricsServer: metricsServer,

		RedisOpt:  redisOpt,
		WorkQueue: workQueue,
	}

	return ct, nil
}

// Container provides common dependencies for the app, and handles their lifecycle.
type Container struct {
	Cfg      *config.Config
	Log      *zap.Logger
	LogLevel zap.AtomicLevel

	MeterProvider  *sdkmetric.MeterProvider
	TracerProvider *sdktrace.TracerProvider

	DebugServer   *debug.Server
	MetricsServer *telemetry.MetricsServer

	RedisOpt  asynq.RedisClientOpt
	WorkQueue work.Queue
}

func (ct *Container) Close() {
	if err := ct.TracerProvider.Shutdown(context.Background()); err != nil {
		ct.Log.Error("Error shutting down OTEL tracer provider", zap.Error(err))
	}

	if err := ct.MetricsServer.Stop(context.Background()); err != nil {
		ct.Log.Error("Error stopping metrics server", zap.Error(err))
	}

	if err := ct.DebugServer.Stop(context.Background()); err != nil {
		ct.Log.Error("Error stopping debug server", zap.Error(err))
	}

	telemetry.SyncLog(ct.Log)
}
