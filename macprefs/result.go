package macprefs

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
)

type Result struct {
	Success string
	Err     error
}

func (r Result) CobraUtilResult(cmd cobrautil.Cmd) cobrautil.CmdResult {
	return cobrautil.NewCmdResult(cmd, r.Success, r.Err)
}
