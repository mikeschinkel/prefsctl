package preftemplates

import (
	"errors"
)

var (
	ErrFailedToLoadYAMLFile      = errors.New("failed to load YAML file")
	ErrFailedToUnmarshalYAMLFile = errors.New("failed to unmarshal YAML file")
)
