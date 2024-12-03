package macprefs

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
)

type ApplyArgs struct {
	Filename Filename
}

func Apply(ctx Context, cmd *cobrautil.Command, args ApplyArgs) (result cobrautil.Result, err error) {
	return cobrautil.NoResult, err
}
