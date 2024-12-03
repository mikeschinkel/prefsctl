//go:build test

package slogutil

import (
	"bytes"
	"fmt"
)

var TestLogOutput bytes.Buffer
var LogWriter = &TestLogOutput

func init() {
	fmt.Printf("+++ INFO: Log output being buffered to `slogutil.TestLogOutput` as `-tags=test` is being used.\n")
}
