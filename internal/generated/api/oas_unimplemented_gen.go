// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// GetHealthzLive implements getHealthzLive operation.
//
// Returns the current liveness status.
//
// GET /healthz/live
func (UnimplementedHandler) GetHealthzLive(ctx context.Context) (r GetHealthzLiveRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetHealthzReady implements getHealthzReady operation.
//
// Returns the current readiness status.
//
// GET /healthz/ready
func (UnimplementedHandler) GetHealthzReady(ctx context.Context) (r GetHealthzReadyRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}