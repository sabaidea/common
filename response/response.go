package response

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type RespondWithErrorParams struct {
	CTX        context.Context
	Writer     http.ResponseWriter
	StatusCode int
	Message    string
	Error      error
	TraceID    string
}

type RespondWithSuccessParams struct {
	CTX        context.Context
	Writer     http.ResponseWriter
	StatusCode int
	Data       interface{}
}

type ErrResponse struct {
	Code      int    `json:"code"`
	Reason    string `json:"reason"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
	Timestamp string `json:"timestamp"`
}

type ErrResponseParams struct {
	Code      int
	Error     error
	Message   string
	ErrorCode string
}

// NewErrorResponse constructs a standardized error response.
func NewErrorResponse(p ErrResponseParams) ErrResponse {
	reason := "internal server error"
	if p.Error != nil {
		reason = p.Error.Error()
	}
	return ErrResponse{
		Code:      p.Code,
		Reason:    reason,
		Message:   p.Message,
		ErrorCode: p.ErrorCode,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

// RespondWithError sends a standardized JSON error response.
func RespondWithError(p RespondWithErrorParams) {
	if err := p.validate(); err != nil {
		http.Error(p.Writer, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(p.StatusCode)

	response := NewErrorResponse(ErrResponseParams{
		Code:      p.StatusCode,
		Error:     p.Error,
		Message:   p.Message,
		ErrorCode: p.TraceID,
	})

	if err := json.NewEncoder(p.Writer).Encode(map[string]interface{}{
		"error": response,
	}); err != nil {
		http.Error(p.Writer, "failed to encode response", http.StatusInternalServerError)
	}
}

// RespondWithSuccess sends a standardized JSON success response.
func RespondWithSuccess(p RespondWithSuccessParams) {
	if err := p.validate(); err != nil {
		http.Error(p.Writer, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	p.Writer.Header().Set("Content-Type", "application/json")
	p.Writer.WriteHeader(p.StatusCode)

	if err := json.NewEncoder(p.Writer).Encode(p.Data); err != nil {
		http.Error(p.Writer, "failed to encode response", http.StatusInternalServerError)
	}
}

func (p RespondWithErrorParams) validate() error {
	if p.Writer == nil {
		return errors.New("writer is required")
	}
	if !(400 <= p.StatusCode && p.StatusCode < 600) {
		return errors.New("status code must be in the range 400-599 for error responses")
	}
	if p.Error == nil {
		return errors.New("error is required")
	}
	if p.TraceID == "" {
		return errors.New("trace ID is required")
	}
	return nil
}

func (p RespondWithSuccessParams) validate() error {
	if p.Writer == nil {
		return errors.New("writer is required")
	}
	if !(200 <= p.StatusCode && p.StatusCode < 300) {
		return errors.New("status code must be in the range 200-299 for success responses")
	}
	return nil
}
