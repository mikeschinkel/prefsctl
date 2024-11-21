package macprefs

import (
	"errors"
)

var (
	ErrInvalidInput                   = errors.New("invalid preference")
	ErrNotFound                       = errors.New("preference not found")
	ErrUnsupportedType                = errors.New("unsupported preference type")
	ErrUnknown                        = errors.New("unknown preference error")
	ErrNotMacOS                       = errors.New("not running on macOS")
	ErrUnrecognizedMacOSVersion       = errors.New("unrecognized macOS version")
	ErrUnrecognizedMacOSVersionFormat = errors.New("unrecognized macOS version format")
	ErrFailedGettingMacOSVersion      = errors.New("failed to get macOS version")
	ErrFailedGettingMacOSVersionName  = errors.New("failed to get macOS version name")
	ErrFailedParsingMajorVersion      = errors.New("failed parsing major version")
)
