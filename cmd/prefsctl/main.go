package main

import (
	"os"

	"github.com/mikeschinkel/prefsctl/cmd/prefsctl/cmds"
	"github.com/mikeschinkel/prefsctl/cobrautil"
	_ "github.com/mikeschinkel/prefsctl/macprefs/prefdefaults"
)

func main() {
	_, err := cobrautil.Execute(cmds.RootCmd, cobrautil.ExecuteArgs{})
	if err != nil {
		os.Exit(1)
	}
}
