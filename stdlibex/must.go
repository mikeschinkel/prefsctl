package stdlibex

import (
	"io"
)

func MustClose(c io.Closer) {
	err := c.Close()
	if err != nil {
		panic(err)
	}
}

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
