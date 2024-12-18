// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int64 {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int64) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// GetHealthzLiveNoContent is response for GetHealthzLive operation.
type GetHealthzLiveNoContent struct{}

func (*GetHealthzLiveNoContent) getHealthzLiveRes() {}

// GetHealthzReadyNoContent is response for GetHealthzReady operation.
type GetHealthzReadyNoContent struct{}

func (*GetHealthzReadyNoContent) getHealthzReadyRes() {}

// Ref: #/components/schemas/HealthStatus
type HealthStatus struct {
	Errors []string `json:"errors"`
}

// GetErrors returns the value of Errors.
func (s *HealthStatus) GetErrors() []string {
	return s.Errors
}

// SetErrors sets the value of Errors.
func (s *HealthStatus) SetErrors(val []string) {
	s.Errors = val
}

func (*HealthStatus) getHealthzLiveRes()  {}
func (*HealthStatus) getHealthzReadyRes() {}
