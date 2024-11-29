package app

import (
	"github.com/hibiken/asynq"

	"github.com/ianunruh/go-backend-app/internal/work"
)

func (ct *Container) NewAsynqServer() *asynq.Server {
	return asynq.NewServer(ct.RedisOpt, work.AsynqConfig(ct.Cfg.Work, ct.Log))
}

func (ct *Container) NewAsynqServeMux() (*asynq.ServeMux, error) {
	mux := work.NewServeMux(ct.MeterProvider, ct.TracerProvider, ct.Log)

	// TODO register task handlers

	return mux, nil
}

func (ct *Container) NewAsynqScheduler() (*asynq.Scheduler, error) {
	sch := work.NewAsynqScheduler(ct.RedisOpt, ct.Log)

	// TODO register scheduled tasks

	return sch, nil
}
