package macprefs

import (
	"github.com/mikeschinkel/prefsctl/cliutil"
)

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, args ApplyArgs) (result cliutil.Result, err error) {
	return cliutil.NoResult, err
}
