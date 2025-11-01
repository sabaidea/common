package httpresponse

import (
	"encoding/json"
	"net/http"

	"github.com/sabaidea/common/v2/apperror"
	"github.com/sabaidea/common/v2/utils"
)

// SuccessResponse represents the standard success response format
type SuccessResponse struct {
	TraceID string      `json:"trace_id"`          // Correlation/Trace ID for debugging and logging
	Data    interface{} `json:"data"`              // The actual response payload
	Message string      `json:"message,omitempty"` // Optional message (human-readable)
}

// ErrorResponse represents the standard error response format
type ErrorResponse struct {
	TraceID string `json:"trace_id"` // Correlation/Trace ID for debugging and logging
	Error   Error  `json:"error"`    // Structured error object
}

// Error provides detailed error information
type Error struct {
	Reason  string `json:"reason"`            // Technical reason (usually the error message)
	Message string `json:"message,omitempty"` // Human-readable error message
}

// JSONSuccess sends a standardized JSON success response
func JSONSuccess(w http.ResponseWriter, r *http.Request, data interface{}, msg ...string) {
	status, resp := buildSuccess(r, data, msg...)
	writeJSON(w, status, map[string]interface{}{"response": resp})
}

// JSONError sends a standardized JSON error response
func JSONError(w http.ResponseWriter, r *http.Request, err error, msg string) {
	status, resp := buildError(r, err, msg)
	writeJSON(w, status, map[string]interface{}{"response": resp})
}

// --- helpers ---

// buildSuccess constructs a success response with trace ID and optional message
func buildSuccess(r *http.Request, data interface{}, msg ...string) (int, SuccessResponse) {
	traceID := utils.GetTraceID(r.Context())

	resp := SuccessResponse{
		TraceID: traceID,
		Data:    data,
	}
	if len(msg) > 0 {
		resp.Message = msg[0]
	}
	return http.StatusOK, resp
}

// buildError constructs an error response mapped to an HTTP status code
func buildError(r *http.Request, err error, msg string) (int, ErrorResponse) {
	traceID := utils.GetTraceID(r.Context())
	status := apperror.HTTPStatus(err)

	resp := ErrorResponse{
		TraceID: traceID,
		Error: Error{
			Reason:  err.Error(), // The raw error (technical reason)
			Message: msg,         // Custom error message for the client
		},
	}

	return status, resp
}

// writeJSON writes the final JSON response with the given status code
func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
