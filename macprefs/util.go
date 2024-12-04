package macprefs

import (
	"fmt"
	"io"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}

func writeString(w io.Writer, s string) {
	_, _ = w.Write([]byte(s))
}
func writeByte(w io.Writer, b byte) {
	_, _ = w.Write([]byte{b})
}

func noop(any) {}
