package debug

import (
	"context"
	"errors"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	"go.uber.org/zap"
)

func NewServer(cfg Config, log *zap.Logger, logLevel zap.AtomicLevel) *Server {
	mux := newMux(logLevel)

	httpSrv := &http.Server{
		Handler: mux,
	}

	return &Server{
		cfg:     cfg,
		httpSrv: httpSrv,
		log:     log,
	}
}

func newMux(logLevel zap.AtomicLevel) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/debug/log/level", logLevel)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	return mux
}

type Server struct {
	cfg     Config
	httpSrv *http.Server
	log     *zap.Logger
}

func (s *Server) Start() error {
	if s.cfg.ListenAddr == "" {
		return nil
	}

	s.log.Info("Starting debug server", zap.String("addr", s.cfg.ListenAddr))

	ln, err := net.Listen("tcp", s.cfg.ListenAddr)
	if err != nil {
		return err
	}

	go func() {
		if err := s.httpSrv.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.log.Error("Error serving debug HTTP", zap.Error(err))
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return s.httpSrv.Shutdown(ctx)
}
