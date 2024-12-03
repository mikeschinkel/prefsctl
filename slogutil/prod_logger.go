//go:build !test

package slogutil

import (
	"os"
)

var LogWriter = os.Stdout
