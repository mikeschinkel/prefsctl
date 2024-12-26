package cmds

import (
	"github.com/mikeschinkel/prefsctl/cobrautil"
	"github.com/spf13/cobra"
)

func init() {
	cobrautil.AddInitializer(func(cli *CLI) {
		RootCmd.AddCmd(getCmd)
	})
}

var getProps = &GetProps{}

type GetProps struct {
	BaseProps
	//output macprefs.OutputPtr
	//filename macprefs.FilenamePtr
}

var getCmd = NewCmdFromOpts(CmdOpts{
	Parent: RootCmd,
	Command: &cobra.Command{
		Use:   "get",
		Short: "Get preferences",
	},
	Props: getProps,
	Flags: []*CmdFlag{
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
	//RunFunc: runGetFunc,
})

//func runGetFunc(ctx Context, cmd Cmd) error {
//	//p := cmd.Props.(*GetProps)
//	return macprefs.Get(ctx, cmd, macprefs.GetArgs{
//		//Output: macprefs.OutputFormat(*p.output),
//	})
//}
