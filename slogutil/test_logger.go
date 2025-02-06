//go:build test

package slogutil

import (
	"bytes"
	"fmt"
	"log/slog"
)

var TestLogOutput bytes.Buffer
var LogWriter = &TestLogOutput

func init() {
	fmt.Printf("+++ INFO: Log output being buffered to `slogutil.TestLogOutput` as `-tags=test` is being used.\n")
}

func PanicInTest(_ *slog.Logger, msg string, args ...any) error {
	panic(formatForErr(msg, args...))
	return nil
}
