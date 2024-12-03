package slogutil

import (
	"errors"
	"fmt"

	"github.com/mikeschinkel/prefsctl/errutil"
	"github.com/mikeschinkel/prefsctl/macprefs/logargs"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

func AppendLogArgs(args ...any) []any {
	args, _ = appendLogArgs(args...)
	return args
}

func appendLogArgs(args ...any) (_ []any, err error) {
	i := 0
	newArgs := make([]any, 0, len(args)+1)
	errs := errutil.MultiErr{}
	for i < len(args) {
		if _args, ok := args[i].([]any); ok {
			newArgs, err = appendLogArgs(newArgs, _args)
			errs.Add(err)
			i++
			continue
		}
		if i == len(args)-1 {
			funcName := stdlibex.Caller(2).Function
			err = errors.Join(ErrNotEnoughArgsPassed, fmt.Errorf("%s=%s", logargs.CallerLogArg, funcName))
			packageSlog.Debug("not enough parameters", logargs.CallerLogArg, funcName, logargs.ErrorLogArg, err)
			args = append(args, []any{"{!INVALID}"})
		}
		newArgs = append(newArgs, args...)
	}
	err = errs.Err()
	if err != nil {
		packageSlog.Error("missing label or value passed to AppendLogArgs", logargs.ErrorLogArg, err)
	}
	return newArgs, err
}
