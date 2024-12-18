package errutil

import (
	"strings"
)

func CleanErrString(err error) string {
	return strings.Replace(err.Error(), "\n", "; ", -1)
}
