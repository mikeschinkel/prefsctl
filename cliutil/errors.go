package cliutil

import (
	"errors"
	"reflect"

	"github.com/mikeschinkel/prefsctl/errutil"
)

var (
	ErrNoHomeDirVar         = errors.New("neither $XDG_CONFIG_HOME nor $HOME are defined")
	ErrCancelledWithCtrlC   = errors.New("cancelled with ctrl+C")
	ErrNotExpectedInterface = errors.New("not the interface we expected")
)

func NotExpectedInterface(expectedType string, receivedType any) error {
	return errutil.AnnotateErr(
		ErrNotExpectedInterface,
		"expected_type=%s&received_type=%s",
		expectedType,
		reflect.TypeOf(receivedType).String(),
	)
}
