package cmds

import (
	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	getCmd.AddCommand(getDefaultsCmd)
}

var getDefaultsProps = &GetDefaultsProps{}

type GetDefaultsProps struct {
	cliutil.BaseProps
	//filename macprefs.FilenamePtr
	//dummy *string
}

var getDefaultsCmd = cliutil.NewCommandFromArgs(cliutil.CommandOpts{
	Parent: getCmd,
	Command: &cobra.Command{
		Use:   "defaults",
		Short: "Get preference defaults",
	},
	Props: getDefaultsProps,
	Flags: []*cliutil.CommandFlag{
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
	RunFunc: func(ctx Context, cmd *cliutil.Command, props cliutil.Props) (cliutil.Result, error) {
		//p := props.(*GetDefaultsProps)
		return macprefs.GetDefaults(ctx, cmd, macprefs.GetDefaultsArgs{
			//Dummy: *p.dummy,
		})
	},
})
