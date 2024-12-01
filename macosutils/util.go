package macosutils

import (
	"fmt"
)

func panicf(format string, args ...any) {
	panic(fmt.Sprintf(format, args...))
}
