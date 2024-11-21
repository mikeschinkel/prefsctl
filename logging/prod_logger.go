//go:build !test

package logging

import (
	"os"
)

var LogWriter = os.Stdout
