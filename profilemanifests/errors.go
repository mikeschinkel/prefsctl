package profilemanifests

import (
	"errors"
)

var (
	ErrFailedToOpenManifestFile     = errors.New("failed to open manifest file")
	ErrFailedToCreateManifestFile   = errors.New("failed to create manifest file")
	ErrFailedToSaveManifestFile     = errors.New("failed to save manifest file")
	ErrFailedToDecodeManifest       = errors.New("failed to decode manifest")
	ErrFailedToEncodeManifest       = errors.New("failed to encode manifest")
	ErrKeyNotFoundInManifest        = errors.New("key not found")
	ErrMissingPreferenceDomain      = errors.New("missing preference domain")
	ErrInvalidManifestFormatVersion = errors.New("invalid manifest format version")
	ErrKeyNotBooleanType            = errors.New("key is not a boolean type")
	ErrKeyNotIntegerType            = errors.New("key is not a integer type")
	ErrKeyNotStringType             = errors.New("key is not a string type")
)
