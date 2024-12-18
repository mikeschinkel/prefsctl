package macosutils

import (
	"errors"
)

var (
	ErrFailedParsingMajorVersion      = errors.New("failed parsing major version")
	ErrFailedToGetMacOSVersion        = errors.New("failed to get macOS version")
	ErrFailedToGetMacOSVersionName    = errors.New("failed to get macOS version name")
	ErrNotMacOS                       = errors.New("not running on macOS")
	ErrUnrecognizedMacOSVersion       = errors.New("unrecognized macOS version")
	ErrUnrecognizedMacOSVersionFormat = errors.New("unrecognized macOS version format")

	ErrUnknownPreferenceError = errors.New("unknown preference error")
	ErrInvalidInput           = errors.New("invalid preference")
	ErrUnsupportedType        = errors.New("unsupported preference type")
	ErrPreferenceNotFound      = errors.New("preference not found")

	ErrFailedToCreateCFString = errors.New("failed to create CFString")
	ErrFailedToGetPrefNames   = errors.New("failed to get preference names")
)
