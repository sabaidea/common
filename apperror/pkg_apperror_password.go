package apperror

import "errors"

var (
	ErrEmptyPassword    = errors.New("empty password")
	ErrWeakPassword     = errors.New("weak password")
	ErrTooLargePassword = errors.New("too large password")
)
