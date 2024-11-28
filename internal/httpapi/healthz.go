package httpapi

import (
	"context"

	"github.com/ianunruh/go-backend-app/internal/generated/api"
)

func (h *Handlers) GetHealthzLive(ctx context.Context) (api.GetHealthzLiveRes, error) {
	return &api.GetHealthzLiveNoContent{}, nil
}

func (h *Handlers) GetHealthzReady(ctx context.Context) (api.GetHealthzReadyRes, error) {
	return &api.GetHealthzReadyNoContent{}, nil
}
