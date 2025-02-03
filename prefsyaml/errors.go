package prefsyaml

import (
	"errors"
)

var (
	ErrFailedToDecodeYAML    = errors.New("failed to decode YAML")
	ErrFailedToMarshalToYAML = errors.New("failed to marshal to YAML")
)
