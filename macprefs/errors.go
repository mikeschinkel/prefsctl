package macprefs

import (
	"errors"
)

var (
	ErrFailedToQueryPrefState     = errors.New("failed to query preference state")
	ErrFailedToQueryGroups        = errors.New("failed to query groups")
	ErrUnsupportedType            = errors.New("unsupported preference type")
	ErrUnsupportedMacOSVersion    = errors.New("unsupported macOS version")
	ErrUnexpectedPreferenceType   = errors.New("unexpected preference type")
	ErrInvalidOutputFormat        = errors.New("not a valid output format")
	ErrNoPrefsFoundInResourceSpec = errors.New("no prefs found in resource spec")
)
