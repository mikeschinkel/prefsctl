package cmds

import (
	//"reflect"

	"github.com/mikeschinkel/prefsctl/cliutil"
	"github.com/mikeschinkel/prefsctl/macprefs"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getProps = &GetProps{}

type GetProps struct {
	cliutil.BaseProps
	//output macprefs.OutputPtr
	//filename macprefs.FilenamePtr
}

var getCmd = cliutil.NewCommandFromArgs(cliutil.CommandOpts{
	Parent: rootCmd,
	Command: &cobra.Command{
		Use:   "get",
		Short: "Get preferences",
	},
	Props: getProps,
	Flags: []*cliutil.CommandFlag{
		//{
		//	Name:      macprefs.OutputFlag,
		//	Type:      reflect.String,
		//	Required:  false,
		//	Shorthand: 'o',
		//	AssignFunc: func(value any) {
		//		getProps.output = macprefs.OutputPtr(value.(*string))
		//	},
		//},
		//{
		//	Name:      macprefs.FilenameFlag,
		//	Type:      reflect.String,
		//	Required:  true,
		//	Shorthand: 'f',
		//	AssignFunc: func(value any) {
		//		getProps.filename = macprefs.FilenamePtr(value.(*string))
		//	},
		//},
	},
	RunFunc: func(ctx Context, cmd *cliutil.Command, props cliutil.Props) (cliutil.Result, error) {
		//p := props.(*GetProps)
		return macprefs.Get(ctx, cmd, macprefs.GetArgs{
			//Output: macprefs.OutputFormat(*p.output),
		})
	},
})
