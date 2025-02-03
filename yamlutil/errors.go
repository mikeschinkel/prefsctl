package yamlutil

import (
	"errors"
)

var (
	ErrFailedToYAMLEncodeObject = errors.New("failed to YAML encode object")
	ErrFailedToCloseYAMLEncoder = errors.New("failed to close YAML encoder")
	ErrFailedToOpenYAMLFile     = errors.New("failed to open YAML file")
)
