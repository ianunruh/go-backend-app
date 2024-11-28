package telemetry

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.uber.org/zap"
)

type MetricsConfig struct {
	ListenAddr string `yaml:"listenAddr" env:"LISTEN_ADDR"`
}

func NewOTELMeterProvider() (*sdkmetric.MeterProvider, error) {
	reader, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))

	return provider, nil
}

func NewMetricsServer(cfg MetricsConfig, log *zap.Logger) *MetricsServer {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	httpSrv := &http.Server{
		Handler: mux,
	}

	return &MetricsServer{
		cfg:     cfg,
		httpSrv: httpSrv,
		log:     log,
	}
}

type MetricsServer struct {
	cfg     MetricsConfig
	httpSrv *http.Server
	log     *zap.Logger
}

func (s *MetricsServer) Start() error {
	if s.cfg.ListenAddr == "" {
		return nil
	}

	s.log.Info("Starting Prometheus metrics server", zap.String("addr", s.cfg.ListenAddr))

	ln, err := net.Listen("tcp", s.cfg.ListenAddr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.httpSrv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("Error serving metrics HTTP", zap.Error(err))
		}
	}()

	return nil
}

func (s *MetricsServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return s.httpSrv.Shutdown(ctx)
}
