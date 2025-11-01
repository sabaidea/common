package utils

import (
	"context"
	"fmt"

	"github.com/sabaidea/common/v2/apperror"
	"github.com/sabaidea/common/v2/constant"
	"github.com/sabaidea/common/v2/logger"
)

// GetTraceID retrieves the trace ID from the context.
// It returns the trace ID if found, otherwise logs an error and returns an empty string.
func GetTraceID(ctx context.Context) string {
	if val := ctx.Value(constant.TraceIDKey); val != nil {
		if traceID, ok := val.(string); ok {
			return traceID
		}
	}
	logger.GetLogger().Error(&logger.Log{
		Event:      "get trace id",
		Error:      fmt.Errorf("%w: trace id not found in context", apperror.ErrInvalidInput),
		TraceID:    "config",
		Additional: map[string]interface{}{"key": constant.TraceIDKey},
	})
	return ""
}
