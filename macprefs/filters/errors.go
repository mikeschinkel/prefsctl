package filters

import (
	"errors"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	ErrUnknown                      = errors.New("unknown")
	ErrFailedToAccessGroupKeyValues = errors.New("failed to access key values for group")
	ErrFailedToInitializeGroup      = errors.New("failed to initialize group")
	ErrFailedTypeAssertion          = errors.New("invalid type assertion")
	ErrInvalidRegExp                = errors.New("invalid regexp")
	ErrInvalidFilterMatchCriteria   = errors.New("invalid filter match criteria")
)
