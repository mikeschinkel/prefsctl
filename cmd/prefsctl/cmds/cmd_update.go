package cmds

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/mikeschinkel/prefsctl/macprefs"
)

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {
		RootCmd.AddCmd(updateCmd)
	})
}

var updateProps = &UpdateProps{}

type UpdateProps struct {
	BaseProps
	//filename macprefs.FilenamePtr
}

// updateCmd represents the `ip deallocate` command
var updateCmd = NewCmdFromOpts(CmdOpts{
	Parent: RootCmd,
	Command: &Command{
		Use:   "update",
		Short: "Update defaults",
		Long:  "Update the defaults YAML file from https://github.com/ProfileManifests/ProfileManifests",
	},
	Props: updateProps,
	//Flags: []*CmdFlag{
	//	{
	//		Name:      FilenameFlagName,
	//		Type:      reflect.String,
	//		Required:  true,
	//		Shorthand: FilenameFlagShorthand,
	//		AssignFunc: func(value any) {
	//			updateProps.filename = value.(*string)
	//		},
	//	},
	//},
	RunFunc: runUpdateFunc,
})

func runUpdateFunc(ctx Context, cfg cobrautil.Config, cmd Cmd) cobrautil.CmdResult {
	return macprefs.Update(ctx, cfg, cmd, macprefs.UpdateArgs{
		//Filename: macprefs.Filename(*cmd.Props().(*UpdateProps).filename),
	}).CobraUtilResult(cmd)
}
