package cmds

import (
	"context"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/spf13/cobra"
)

var (
	NewCmdFromOpts       = cobrautil.NewCmdFromOpts
	PersistentBoolFlag   = cobrautil.PersistentBoolFlag
	PersistentStringFlag = cobrautil.PersistentStringFlag
)

type (
	Context = context.Context

	BaseProps   = cobrautil.BaseProps
	CLI         = cobrautil.CLI
	Cmd         = cobrautil.Cmd
	CmdFlag     = cobrautil.CmdFlag
	CmdFlagArgs = cobrautil.CmdFlagArgs
	CmdOpts     = cobrautil.CmdOpts

	Command = cobra.Command
)
