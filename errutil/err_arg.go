package errutil

import (
	"fmt"
)

func ErrArg(name ErrKey, value ...any) error {
	switch len(value) {
	case 0:
		panicf("ErrArg requires at least two parameters, called for %s", name)
	case 1:
		// value already a single value; do nothing
	default:
		value[0] = fmt.Sprintf(value[0].(string), value[1:]...)
	}
	return fmt.Errorf("%s=%v", name, value[0])
}
