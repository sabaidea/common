package apperror

import (
	"errors"
	"fmt"
)

// Core application/domain errors
var (
	ErrEmpty          = errors.New("empty")
	ErrDuplicateEntry = errors.New("duplicate entry")
	ErrNotFound       = errors.New("not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrTooMany        = errors.New("too many requests")
	ErrUnavailable    = errors.New("service unavailable")
	ErrTimeout        = errors.New("timeout")
	ErrInternal       = errors.New("internal error")
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
