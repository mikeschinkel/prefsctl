package macprefs

import (
	"errors"
)

var (
	ErrInvalidInput    = errors.New("invalid preference")
	ErrNotFound        = errors.New("preference not found")
	ErrUnsupportedType = errors.New("unsupported preference type")
	ErrUnknown         = errors.New("unknown preference error")
)
