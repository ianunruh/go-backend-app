package httpapi

import (
	"context"
	"net/http"

	"github.com/ianunruh/go-backend-app/internal/generated/api"
)

func NewHandlers() *Handlers {
	return &Handlers{}
}

type Handlers struct {
}

func (h *Handlers) NewError(ctx context.Context, err error) *api.ErrorStatusCode {
	return &api.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: api.Error{
			Message: err.Error(),
		},
	}
}
