// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	ht "github.com/ogen-go/ogen/http"
)

func encodeGetHealthzLiveResponse(response GetHealthzLiveRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetHealthzLiveNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *HealthStatus:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetHealthzReadyResponse(response GetHealthzReadyRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *GetHealthzReadyNoContent:
		w.WriteHeader(204)
		span.SetStatus(codes.Ok, http.StatusText(204))

		return nil

	case *HealthStatus:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(503)
		span.SetStatus(codes.Error, http.StatusText(503))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeErrorResponse(response *ErrorStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := new(jx.Encoder)
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
