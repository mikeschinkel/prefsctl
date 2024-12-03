//go:build test

package logutil

import (
	"bytes"
	"fmt"
)

var TestLogOutput bytes.Buffer
var LogWriter = &TestLogOutput

func init() {
	fmt.Printf("+++ INFO: Log output being buffered to `logutil.TestLogOutput` as `-tags=test` is being used.\n")
}
