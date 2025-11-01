package apperror

import "errors"

var (
	ErrEmptyPhoneNumber   = errors.New("phone: empty")
	ErrInvalidPhoneNumber = errors.New("phone: invalid format")
	ErrTooLongPhoneNumber = errors.New("phone: too long (>15 digits)")
)
