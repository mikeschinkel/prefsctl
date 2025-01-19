package prefsyaml

import (
	"errors"
)

var (
	ErrFoo                = errors.New("foo")
	ErrFailedToDecodeYAML = errors.New("failed to decode YAML")
)
