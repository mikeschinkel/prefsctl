//go:build !test

package slogutil

import (
	"os"
)

var LogWriter = os.Stdout

func PanicInTest(log *SlogLogger, msg string, args ...any) error {
	return formatForErr(msg, args...)
}
