//go:build !test

package logutil

import (
	"os"
)

var LogWriter = os.Stdout
