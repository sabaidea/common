package apperror

import "errors"

var (
	ErrEmptyEmail      = errors.New("email: empty")
	ErrInvalidEmail    = errors.New("email: invalid format")
	ErrLocalTooLong    = errors.New("email: local part too long (>64)")
	ErrDomainTooLong   = errors.New("email: domain too long (>255)")
	ErrMultipleAtSigns = errors.New("email: multiple @ signs")
)
