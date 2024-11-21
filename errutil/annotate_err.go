package errutil

import (
	"errors"
	"fmt"
)

func AnnotateErr(err error, format string, args ...any) error {
	var err2 error
	if len(args) == 0 {
		err2 = errors.New(format)
	} else {
		err2 = fmt.Errorf(format, args...)
	}
	return errors.Join(err, err2)
}
