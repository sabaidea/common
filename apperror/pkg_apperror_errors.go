package apperror

import (
	"errors"
	"fmt"
)

// Core application/domain errors
var (
	ErrEmpty          = errors.New("EMPTY")
	ErrDuplicateEntry = errors.New("DUPLICATE_ENTRY")
	ErrNotFound       = errors.New("NOT_FOUND")
	ErrInvalidInput   = errors.New("INVALID_INPUT")
	ErrUnauthorized   = errors.New("UNAUTHORIZED")
	ErrForbidden      = errors.New("FORBIDDEN")
	ErrTooMany        = errors.New("TOO_MANY_REQUESTS")
	ErrUnavailable    = errors.New("SERVICE_UNAVAILABLE")
	ErrTimeout        = errors.New("TIMEOUT")
	ErrInternal       = errors.New("INTERNAL_ERROR")
)

func NewErrEmpty(msg string) error {
	return fmt.Errorf("%w: %s", ErrEmpty, msg)
}

func NewErrDuplicateEntry(msg string) error {
	return fmt.Errorf("%w: %s", ErrDuplicateEntry, msg)
}

func NewErrNotFound(msg string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, msg)
}

func NewErrInvalidInput(msg string) error {
	return fmt.Errorf("%w: %s", ErrInvalidInput, msg)
}

func NewErrUnauthorized(msg string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, msg)
}

func NewErrForbidden(msg string) error {
	return fmt.Errorf("%w: %s", ErrForbidden, msg)
}

func NewErrTooMany(msg string) error {
	return fmt.Errorf("%w: %s", ErrTooMany, msg)
}

func NewErrUnavailable(msg string) error {
	return fmt.Errorf("%w: %s", ErrUnavailable, msg)
}

func NewErrTimeout(msg string) error {
	return fmt.Errorf("%w: %s", ErrTimeout, msg)
}

func NewErrInternal(msg string) error {
	return fmt.Errorf("%w: %s", ErrInternal, msg)
}
