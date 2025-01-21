package applemdm

import (
	"io"
)

func mustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		slog.Error(err.Error())
	}
}
