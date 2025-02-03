package yamlutil

import (
	"fmt"
	"io"
	"log/slog"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}
