package cmds

import (
	"reflect"

	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {
		RootCmd.AddCmd(applyCmd)
	})
}

var applyProps = &ApplyProps{}

type ApplyProps struct {
	BaseProps
	filename macprefs.FilenamePtr
}

// applyCmd represents the `ip deallocate` command
var applyCmd = NewCmdFromOpts(CmdOpts{
	Parent: RootCmd,
	Command: &Command{
		Use:   "apply -f <filename>",
		Short: "Apply a MacPrefs manifest",
		Long:  "Apply a MacPrefs manifest to set macOS preferences",
	},
	Props: applyProps,
	Flags: []*CmdFlag{
		{
			Name:      FilenameFlagName,
			Type:      reflect.String,
			Required:  true,
			Shorthand: FilenameFlagShorthand,
			AssignFunc: func(value any) {
				applyProps.filename = value.(*string)
			},
		},
	},
	RunFunc: runApplyFunc,
})

func runApplyFunc(ctx Context, cmd Cmd) cobrautil.CmdResult {
	return macprefs.Apply(ctx, cmd, macprefs.ApplyArgs{
		Filename: macprefs.Filename(*cmd.Props().(*ApplyProps).filename),
	}).CobraUtilResult(cmd)
}
