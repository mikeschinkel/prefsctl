package slogutil

import (
	"errors"
	"fmt"
)

var (
	ErrNotEnoughArgsPassed = fmt.Errorf("not enough args passed")
)

func formatForErr(msg string, args ...any) error {
	if len(args) > 0 && len(args)%2 != 0 {
		args = append(args, "%!(MISSING)")
	}
	for i := 0; i < len(args); i += 2 {
		msg += fmt.Sprintf("%v=%v", args[i], args[i+1])
	}
	return errors.New(msg)
}
