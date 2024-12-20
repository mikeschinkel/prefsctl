package cmds

import (
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCmd(getDefaultsCmd)
}

var getDefaultsProps = &GetDefaultsProps{}

type GetDefaultsProps struct {
	BaseProps
	//filename macprefs.FilenamePtr
	//dummy *string
}

var getDefaultsCmd = NewCmdFromOpts(CmdOpts{
	Parent: getCmd,
	Command: &cobra.Command{
		Use:   "defaults",
		Short: "Get preference defaults",
	},
	Props: getDefaultsProps,
	Flags: []*CmdFlag{
		//{
		//	Name:      macprefs.FilenameFlag,
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'f',
		//	AssignFunc: func(value any) {
		//		getDefaultsProps.filename = macprefs.FilenamePtr(value.(*string))
		//	},
		//},
		//{
		//	Name:      "Dummy",
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'd',
		//	AssignFunc: func(value any) {
		//		getDefaultsProps.dummy = value.(*string)
		//	},
		//},
	},
	RunFunc: runGetDefaultsFunc,
})

func runGetDefaultsFunc(ctx Context, cmd Cmd) error {
	//p := cmd.Props.(*GetDefaultsProps)
	return macprefs.GetDefaults(ctx, macprefs.GetDefaultsArgs{
		Printer: cmd,
	})
}
