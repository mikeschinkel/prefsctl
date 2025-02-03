package fsutil

import (
	"io"
	"log/slog"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}
