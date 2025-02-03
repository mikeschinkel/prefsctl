package errutil

import (
	"errors"
)

var (
	ErrBytesReadNotSameAsFileSize = errors.New("bytes read not same as file size")
	ErrGettingFileInfo            = errors.New("error getting file info")
	ErrReadingFile                = errors.New("error reading file")
	ErrOpeningFile                = errors.New("error reading file")
)
