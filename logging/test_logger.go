//go:build test

package logging

import (
	"bytes"
	"fmt"
)

var TestLogOutput bytes.Buffer
var LogWriter = &TestLogOutput

func init() {
	fmt.Printf("+++ INFO: Log output being buffered to `logging.TestLogOutput` as `-tags=test` is being used.\n")
}
