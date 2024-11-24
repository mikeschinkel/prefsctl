package macprefs

import (
	"errors"
)

var (
	ErrFailedParsingMajorVersion      = errors.New("failed parsing major version")
	ErrFailedToCreateCFString         = errors.New("failed to create CFString")
	ErrFailedToGetMacOSVersion        = errors.New("failed to get macOS version")
	ErrFailedToGetMacOSVersionName    = errors.New("failed to get macOS version name")
	ErrFailedToGetPrefDomains         = errors.New("failed to get preference domains")
	ErrFailedToGetDomainPrefs         = errors.New("failed to get domain preferences")
	ErrFailedToQueryPrefState         = errors.New("failed to query preference state")
	ErrFailedToGetPrefNames           = errors.New("failed to get preference names")
	ErrInvalidInput                   = errors.New("invalid preference")
	ErrNotFound                       = errors.New("preference not found")
	ErrNotMacOS                       = errors.New("not running on macOS")
	ErrUnknown                        = errors.New("unknown preference error")
	ErrUnrecognizedMacOSVersion       = errors.New("unrecognized macOS version")
	ErrUnrecognizedMacOSVersionFormat = errors.New("unrecognized macOS version format")
	ErrUnsupportedType                = errors.New("unsupported preference type")
)
